package harmony

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	committypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	hmylctypes "github.com/datachainlab/ibc-harmony-client/modules/light-clients/harmony/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/harmony-one/harmony/block"
	v3 "github.com/harmony-one/harmony/block/v3"
	hmytypes "github.com/harmony-one/harmony/core/types"
	rpcv2 "github.com/harmony-one/harmony/rpc/v2"
	"github.com/harmony-one/harmony/shard"
	"github.com/hyperledger-labs/yui-relayer/core"
)

type Prover struct {
	chain        *Chain  // target shard
	beaconClient *Client // beacon
	config       ProverConfig
}

var _ core.ProverI = (*Prover)(nil)

func NewProver(chain *Chain, config ProverConfig) (*Prover, error) {
	beaconClient := NewHarmonyClient(chain.config.BeaconRpcAddr)
	return &Prover{
		chain:        chain,
		beaconClient: beaconClient,
		config:       config,
	}, nil
}

// GetChainID returns the chain ID
func (pr *Prover) GetChainID() string {
	return pr.chain.ChainID()
}

// QueryLatestHeader returns the latest header from the chain
func (pr *Prover) QueryLatestHeader() (out core.HeaderI, err error) {
	if pr.chain.config.ShardId == 0 {
		return pr.queryLatestHeaderForBeacon()
	} else {
		return pr.queryLatestHeaderForShard()
	}
}

// GetLatestLightHeight returns the latest height on the light client
func (pr *Prover) GetLatestLightHeight() (int64, error) {
	return -1, nil
}

// CreateMsgCreateClient creates a CreateClientMsg to this chain
func (pr *Prover) CreateMsgCreateClient(clientID string, dstHeader core.HeaderI, signer sdk.AccAddress) (*clienttypes.MsgCreateClient, error) {
	h, ok := dstHeader.(*hmylctypes.Header)
	if !ok {
		return nil, errors.New("dstHeader must be an harmony header")
	}
	beaconHeader, err := decodeV3(h.BeaconHeader.Header)
	if err != nil {
		return nil, err
	}
	if l := len(beaconHeader.ShardState()); l == 0 {
		// The last beacon header of the previous epoch is needed
		prevHeader, err := pr.queryEpochLastHeader(beaconHeader.Epoch().Uint64()-1, true)
		if err != nil {
			return nil, err
		}
		h.EpochHeaders = append([]hmylctypes.BeaconHeader{*prevHeader.BeaconHeader}, h.EpochHeaders...)
		beaconHeader, err = decodeV3(prevHeader.BeaconHeader.Header)
		if err != nil {
			return nil, err
		}
	}

	var targetHeader *v3.Header
	if pr.chain.config.ShardId > 0 {
		var err error
		targetHeader, err = decodeV3(h.ShardHeader)
		if err != nil {
			return nil, err
		}
	} else {
		targetHeader = beaconHeader
	}
	var shardState shard.State
	if err := rlp.DecodeBytes(beaconHeader.ShardState(), &shardState); err != nil {
		return nil, err
	}
	committee, err := shardState.FindCommitteeByID(0)
	if err != nil {
		return nil, err
	}
	committeeRlp, err := rlp.EncodeToBytes(committee)
	if err != nil {
		return nil, err
	}

	tp, err := pr.config.TrustingPeriodDuration()
	if err != nil {
		return nil, err
	}

	clientState := &hmylctypes.ClientState{
		ShardId:         pr.chain.config.ShardId,
		ContractAddress: pr.chain.config.IBCHostAddress().Bytes(),
		LatestEpoch:     targetHeader.Epoch().Uint64(),
		LatestCommittee: committeeRlp,
		LatestHeight:    clienttypes.Height{RevisionNumber: 0, RevisionHeight: targetHeader.Number().Uint64()},
		TrustingPeriod:  tp,
		Frozen:          false,
	}

	proof, err := decodeRLP(h.AccountProof)
	if err != nil {
		return nil, err
	}
	accountRLP, err := hmylctypes.VerifyProof(targetHeader.Root(), pr.chain.config.IBCHostAddress().Bytes(), proof)
	if err != nil {
		return nil, err
	}
	storageHash, err := decodeStorageHash(accountRLP)
	if err != nil {
		return nil, err
	}

	consensusState := &hmylctypes.ConsensusState{
		Timestamp: targetHeader.Time().Uint64(),
		Root:      committypes.NewMerkleRoot(storageHash).Hash,
	}
	return clienttypes.NewMsgCreateClient(
		clientState,
		consensusState,
		signer.String(),
	)
}

// SetupHeader creates a new header based on a given header
func (pr *Prover) SetupHeader(dstChain core.LightClientIBCQueryierI, baseSrcHeader core.HeaderI) (core.HeaderI, error) {
	srcChain := pr.chain
	header, ok := baseSrcHeader.(*hmylctypes.Header)
	if !ok {
		return nil, errors.New("invalid header type")
	}
	dsth, err := dstChain.GetLatestLightHeight()
	if err != nil {
		return nil, err
	}

	// retrieve counterparty client from dst chain
	counterpartyClientRes, err := dstChain.QueryClientState(dsth)
	if err != nil {
		return nil, err
	}

	var cs exported.ClientState
	if err := srcChain.codec.UnpackAny(counterpartyClientRes.ClientState, &cs); err != nil {
		return nil, err
	}
	trustedHeight := cs.GetLatestHeight().GetRevisionHeight()
	trustedHeader, err := srcChain.client.FullHeader(context.Background(), trustedHeight)
	if err != nil {
		return nil, err
	}
	beaconEpoch := header.GetBeaconEpoch()
	gapEpochSize := new(big.Int).Sub(beaconEpoch, trustedHeader.Epoch).Int64()
	if gapEpochSize > 0 {
		epochHeaders := make([]hmylctypes.BeaconHeader, gapEpochSize)
		prevEpoch := trustedHeader.Epoch.Uint64()
		for i := 0; i < int(gapEpochSize); i++ {
			// The last beacon header of epoch is needed to update committee
			prevHeader, err := pr.queryEpochLastHeader(prevEpoch, true)
			if err != nil {
				return nil, err
			}
			epochHeaders[i] = *prevHeader.BeaconHeader
			prevEpoch += 1
		}
		header.EpochHeaders = epochHeaders
	}
	return header, nil
}

// UpdateLightWithHeader updates a header on the light client and returns the header and height corresponding to the chain
func (pr *Prover) UpdateLightWithHeader() (header core.HeaderI, provableHeight int64, queryableHeight int64, err error) {
	h, err := pr.QueryLatestHeader()
	if err != nil {
		return nil, -1, -1, err
	}
	height := int64(h.GetHeight().GetRevisionHeight())
	if err != nil {
		return nil, -1, -1, err
	}
	return h, height, height, nil
}

// QueryClientConsensusState returns the ClientConsensusState and its proof
func (pr *Prover) QueryClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	res, err := pr.chain.QueryClientConsensusState(height, dstClientConsHeight)
	if err != nil {
		return nil, err
	}

	key, err := hmylctypes.ConsensusStateCommitmentSlot(pr.chain.Path().ClientID, dstClientConsHeight)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getStorageProof(hexKey(key), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	res.Proof = proof
	res.ProofHeight = clienttypes.NewHeight(0, uint64(height))
	return res, nil
}

// QueryClientStateWithProof returns the ClientState and its proof
func (pr *Prover) QueryClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	res, err := pr.chain.QueryClientState(height)
	if err != nil {
		return nil, err
	}
	key, err := hmylctypes.ClientStateCommitmentSlot(pr.chain.Path().ClientID)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getStorageProof(hexKey(key), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	res.Proof = proof
	res.ProofHeight = clienttypes.NewHeight(0, uint64(height))
	return res, nil
}

// QueryConnectionWithProof returns the Connection and its proof
func (pr *Prover) QueryConnectionWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	res, err := pr.chain.QueryConnection(height)
	if err != nil {
		return nil, err
	}
	key, err := hmylctypes.ConnectionCommitmentSlot(pr.chain.Path().ConnectionID)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getStorageProof(hexKey(key), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	res.Proof = proof
	res.ProofHeight = clienttypes.NewHeight(0, uint64(height))
	return res, nil
}

// QueryChannelWithProof returns the Channel and its proof
func (pr *Prover) QueryChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	res, err := pr.chain.QueryChannel(height)
	if err != nil {
		return nil, err
	}
	path := pr.chain.Path()
	key, err := hmylctypes.ChannelCommitmentSlot(path.PortID, path.ChannelID)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getStorageProof(hexKey(key), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	res.Proof = proof
	res.ProofHeight = clienttypes.NewHeight(0, uint64(height))
	return res, nil
}

// QueryPacketCommitmentWithProof returns the packet commitment and its proof
func (pr *Prover) QueryPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	res, err := pr.chain.QueryPacketCommitment(height, seq)
	if err != nil {
		return nil, err
	}
	path := pr.chain.Path()
	key, err := hmylctypes.PacketCommitmentSlot(path.PortID, path.ChannelID, seq)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getStorageProof(hexKey(key), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	res.Proof = proof
	res.ProofHeight = clienttypes.NewHeight(0, uint64(height))
	return res, nil
}

// QueryPacketAcknowledgementCommitmentWithProof returns the packet acknowledgement commitment and its proof
func (pr *Prover) QueryPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	res, err := pr.chain.QueryPacketAcknowledgementCommitment(height, seq)
	if err != nil {
		return nil, err
	}
	path := pr.chain.Path()
	key, err := hmylctypes.PacketAcknowledgementCommitmentSlot(path.PortID, path.ChannelID, seq)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getStorageProof(hexKey(key), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	res.Proof = proof
	res.ProofHeight = clienttypes.NewHeight(0, uint64(height))
	return res, nil
}

func (pr *Prover) getAccountProof(client *Client, key []byte, blockNumber *big.Int) ([]byte, error) {
	ethProof, err := getETHProof(client, pr.chain.config.IBCHostAddress(), key, blockNumber)
	if err != nil {
		return nil, err
	}
	return ethProof.AccountProofRLP, nil
}

func (pr *Prover) getStorageProof(key []byte, blockNumber *big.Int) ([]byte, error) {
	ethProof, err := getETHProof(pr.chain.client, pr.chain.config.IBCHostAddress(), key, blockNumber)
	if err != nil {
		return nil, err
	}
	if len(ethProof.StorageProofRLP) == 0 {
		return nil, errors.New("storage proof is empty")
	}
	return ethProof.StorageProofRLP[0], nil
}

// When skipsShardHeader is true, the return header does not contain a shard header.
func (pr *Prover) queryEpochLastHeader(epoch uint64, skipsShardHeader bool) (*hmylctypes.Header, error) {
	height, err := pr.beaconClient.EpochLastBlockNumber(context.Background(), epoch)
	if err != nil {
		return nil, err
	}
	beaconHeader, err := pr.beaconClient.FullHeader(context.Background(), height)
	if err != nil {
		return nil, err
	}
	nextHeader, err := pr.beaconClient.FullHeader(context.Background(), height+1)
	if err != nil {
		return nil, err
	}

	var shardHeader *rpcv2.BlockHeader = nil
	crossLinkIndex := uint32(0)
	if !skipsShardHeader {
		if len(beaconHeader.CrossLink) > 0 {
			var crossLinks hmytypes.CrossLinks
			if err := rlp.DecodeBytes(beaconHeader.CrossLink, &crossLinks); err != nil {
				return nil, err
			}

			var crossLink *hmytypes.CrossLink
			found := false
			for i, cl := range crossLinks {
				if cl.ShardID() == pr.chain.config.ShardId {
					crossLinkIndex = uint32(i)
					crossLink = &crossLinks[i]
					found = true
					break
				}
			}
			if found {
				shardHeader, err := pr.chain.client.FullHeader(context.Background(), crossLink.BlockNumberF.Uint64())
				if err != nil {
					return nil, err
				}
				shv3, err := convertHeader(shardHeader)
				if err != nil {
					return nil, err
				}
				b := block.Header{Header: shv3}
				if !bytes.Equal(crossLink.HashF.Bytes(), b.Hash().Bytes()) {
					return nil, fmt.Errorf("invalid cross link on beacon block %d, shard block %d. expected: %s, got: %s",
						beaconHeader.Number.Uint64(), shardHeader.Number.Uint64(), crossLink.HashF.Hex(), b.Hash().Hex())
				}
			}
		}
	}

	bhRLP, err := encodeAsV3(beaconHeader)
	if err != nil {
		return nil, err
	}
	var shRLP []byte = nil
	var proof []byte = nil
	if shardHeader != nil {
		shRLP, err = encodeAsV3(shardHeader)
		if err != nil {
			return nil, err
		}
		proof, err = pr.getAccountProof(pr.chain.client, nil, shardHeader.Number)
		if err != nil {
			return nil, err
		}
	}
	return &hmylctypes.Header{
		ShardHeader: shRLP,
		BeaconHeader: &hmylctypes.BeaconHeader{
			Header:       bhRLP,
			CommitSig:    nextHeader.LastCommitSignature,
			CommitBitmap: nextHeader.LastCommitBitmap,
		},
		CrossLinkIndex: uint32(crossLinkIndex),
		AccountProof:   proof,
	}, nil
}

func (pr *Prover) queryLatestHeaderForBeacon() (out core.HeaderI, err error) {
	height, err := pr.beaconClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	beaconHeader, err := pr.beaconClient.FullHeader(context.Background(), height-1)
	if err != nil {
		return nil, err
	}
	nextHeader, err := pr.beaconClient.FullHeader(context.Background(), height)
	if err != nil {
		return nil, err
	}

	bhRLP, err := encodeAsV3(beaconHeader)
	if err != nil {
		return nil, err
	}
	proof, err := pr.getAccountProof(pr.beaconClient, nil, beaconHeader.Number)
	if err != nil {
		return nil, err
	}
	header := &hmylctypes.Header{
		ShardHeader: nil,
		BeaconHeader: &hmylctypes.BeaconHeader{
			Header:       bhRLP,
			CommitSig:    nextHeader.LastCommitSignature,
			CommitBitmap: nextHeader.LastCommitBitmap,
		},
		CrossLinkIndex: 0,
		AccountProof:   proof,
	}
	return header, nil
}

func (pr *Prover) queryLatestHeaderForShard() (out core.HeaderI, err error) {
	height, err := pr.beaconClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	// For getting commitSig and commitBitmap from the next height
	height -= 1
	var header *hmylctypes.Header
	// Find a crosslinked header pair.
	// Decrease the beacon height one by one until it is found.
	// TODO consider iterate limit
	for ; height > 0; height-- {
		beaconHeader, err := pr.beaconClient.FullHeader(context.Background(), height)
		if err != nil {
			return nil, err
		}
		if len(beaconHeader.CrossLink) == 0 {
			continue
		}
		var crossLinks hmytypes.CrossLinks
		if err := rlp.DecodeBytes(beaconHeader.CrossLink, &crossLinks); err != nil {
			return nil, err
		}
		crossLinks.Sort()
		var crossLink *hmytypes.CrossLink
		crossLinkIndex := -1
		// If there can be multiple height cross links for the same shard id, use the latest
		for i := len(crossLinks) - 1; i >= 0; i-- {
			cl := crossLinks[i]
			if cl.ShardID() == pr.chain.config.ShardId {
				crossLinkIndex = i
				crossLink = &cl
				break
			}
		}
		if crossLinkIndex == -1 {
			continue
		}

		shardHeader, err := pr.chain.client.FullHeader(context.Background(), crossLink.BlockNumberF.Uint64())
		if err != nil {
			return nil, err
		}
		shv3, err := convertHeader(shardHeader)
		if err != nil {
			return nil, err
		}
		b := block.Header{Header: shv3}
		if !bytes.Equal(crossLink.HashF.Bytes(), b.Hash().Bytes()) {
			return nil, fmt.Errorf("invalid cross link on beacon block %d, shard block %d. expected: %s, got: %s",
				beaconHeader.Number.Uint64(), shardHeader.Number.Uint64(), crossLink.HashF.Hex(), b.Hash().Hex())
		}
		nextBeaconHeader, err := pr.beaconClient.FullHeader(context.Background(), beaconHeader.Number.Uint64()+1)
		if err != nil {
			return nil, err
		}

		bhRLP, err := encodeAsV3(beaconHeader)
		if err != nil {
			return nil, err
		}
		var shRLP []byte = nil
		var proof []byte = nil
		if shardHeader != nil {
			shRLP, err = encodeAsV3(shardHeader)
			if err != nil {
				return nil, err
			}
			proof, err = pr.getAccountProof(pr.chain.client, nil, shardHeader.Number)
			if err != nil {
				return nil, err
			}
		}
		header = &hmylctypes.Header{
			ShardHeader: shRLP,
			BeaconHeader: &hmylctypes.BeaconHeader{
				Header:       bhRLP,
				CommitSig:    nextBeaconHeader.LastCommitSignature,
				CommitBitmap: nextBeaconHeader.LastCommitBitmap,
			},
			CrossLinkIndex: uint32(crossLinkIndex),
			AccountProof:   proof,
		}
		break
	}
	return header, nil
}

func getETHProof(client *Client, address common.Address, key []byte, blockNumber *big.Int) (*ETHProof, error) {
	var k [][]byte = nil
	if len(key) > 0 {
		k = [][]byte{key}
	}
	proof, err := client.GetETHProof(
		address,
		k,
		blockNumber,
	)
	if err != nil {
		return nil, err
	}
	return proof, nil
}

func encodeAsV3(bh *rpcv2.BlockHeader) ([]byte, error) {
	buf := &bytes.Buffer{}
	v3h, err := convertHeader(bh)
	if err != nil {
		return nil, err
	}
	if err := v3h.EncodeRLP(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeV3(header []byte) (*v3.Header, error) {
	stream := rlp.NewStream(bytes.NewBuffer(header), 0)
	sh := v3.NewHeader()
	if err := sh.DecodeRLP(stream); err != nil {
		return nil, err
	}
	return sh, nil
}

func convertHeader(bh *rpcv2.BlockHeader) (*v3.Header, error) {
	//h := block.Header{Header: v3.NewHeader()}
	h := v3.NewHeader()
	h.SetParentHash(bh.ParentHash)
	h.SetCoinbase(common.HexToAddress(bh.Miner))
	h.SetRoot(bh.StateRoot)
	h.SetTxHash(bh.TransactionsRoot)
	h.SetReceiptHash(bh.ReceiptsRoot)
	h.SetOutgoingReceiptHash(bh.OutgoingReceiptsRoot)
	h.SetIncomingReceiptHash(bh.IncomingReceiptsRoot)
	h.SetBloom(bh.LogsBloom)
	h.SetNumber(bh.Number)
	h.SetGasLimit(bh.GasLimit)
	h.SetGasUsed(bh.GasUsed)
	h.SetTime(bh.Timestamp)
	h.SetExtra(bh.ExtraData)
	h.SetMixDigest(bh.MixHash)
	h.SetViewID(bh.ViewID)
	h.SetEpoch(bh.Epoch)
	h.SetShardID(bh.ShardID)
	var commitSig [96]byte
	copy(commitSig[:], bh.LastCommitSignature)
	h.SetLastCommitSignature(commitSig)
	h.SetLastCommitBitmap(bh.LastCommitBitmap)
	h.SetVrf(bh.Vrf)
	h.SetVdf(bh.Vdf)
	h.SetShardState(bh.ShardState)
	h.SetCrossLinks(bh.CrossLink)
	h.SetSlashes(bh.Slashes)
	return h, nil
}

func hexKey(key []byte) []byte {
	return []byte(strings.Join([]string{"0x", hex.EncodeToString(key[:])}, ""))
}
