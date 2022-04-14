package types

import (
	"time"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	tmclient "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	tmtypes "github.com/tendermint/tendermint/types"
)

var _ exported.Header = (*TmHeader)(nil)

// ConsensusState returns the updated consensus state associated with the header
func (h TmHeader) ConsensusState() *tmclient.ConsensusState {
	return &tmclient.ConsensusState{
		Timestamp:          h.GetTime(),
		Root:               commitmenttypes.NewMerkleRoot(h.SignedHeader.Header.AppHash),
		NextValidatorsHash: h.SignedHeader.Header.NextValidatorsHash,
	}
}

// ClientType defines that the Header is a Tendermint consensus algorithm
func (h TmHeader) ClientType() string {
	return exported.Tendermint
}

// GetHeight returns the current height. It returns 0 if the tendermint
// header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h TmHeader) GetHeight() exported.Height {
	revision := clienttypes.ParseChainID(h.SignedHeader.Header.ChainId)
	return clienttypes.NewHeight(revision, uint64(h.SignedHeader.Header.Height))
}

// GetTime returns the current block timestamp. It returns a zero time if
// the tendermint header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h TmHeader) GetTime() time.Time {
	return h.SignedHeader.Header.Time.Time()
}

//// ValidateBasic calls the SignedHeader ValidateBasic function and checks
//// that validatorsets are not nil.
//// NOTE: TrustedHeight and TrustedValidators may be empty when creating client
//// with MsgCreateClient
func (h TmHeader) ValidateBasic() error {
	th, err := h.Header()
	if err != nil {
		return err
	}
	return th.ValidateBasic()
}

func (h TmHeader) Header() (*tmclient.Header, error) {
	header := &tmclient.Header{
		SignedHeader: h.SignedHeader.SignedHeader().ToProto(),
	}
	if h.TrustedHeight != nil {
		header.TrustedHeight = clienttypes.Height(*h.TrustedHeight)
	}
	var err error
	if h.ValidatorSet != nil {
		header.ValidatorSet, err = h.ValidatorSet.ValidatorSet().ToProto()
		if err != nil {
			return nil, err
		}
	}
	if h.TrustedValidators != nil {
		header.TrustedValidators, err = h.TrustedValidators.ValidatorSet().ToProto()
		if err != nil {
			return nil, err
		}
	}
	return header, nil
}

func (sh SignedHeader) ValidateBasic(chainID string) error {
	return sh.SignedHeader().ValidateBasic(chainID)
}

func (sh SignedHeader) SignedHeader() *tmtypes.SignedHeader {
	return &tmtypes.SignedHeader{
		Header: sh.Header.Header(),
		Commit: sh.Commit.Commit(),
	}
}

func NewLightHeaderFromTm(h *tmtypes.Header) *LightHeader {
	return &LightHeader{
		Version:            (*Consensus)(&h.Version),
		ChainId:            h.ChainID,
		Height:             h.Height,
		Time:               NewTimestampFromTime(h.Time),
		LastBlockId:        NewCanonicalBlockIDFromTm(&h.LastBlockID),
		LastCommitHash:     h.LastCommitHash,
		DataHash:           h.DataHash,
		ValidatorsHash:     h.ValidatorsHash,
		NextValidatorsHash: h.NextValidatorsHash,
		ConsensusHash:      h.ConsensusHash,
		AppHash:            h.AppHash.Bytes(),
		LastResultsHash:    h.LastResultsHash,
		EvidenceHash:       h.EvidenceHash,
		ProposerAddress:    h.ProposerAddress,
	}
}

func NewSignedHeaderFromTm(sh *tmtypes.SignedHeader) *SignedHeader {
	return &SignedHeader{
		Header: NewLightHeaderFromTm(sh.Header),
		Commit: NewCommitFromTm(sh.Commit),
	}
}

func (h LightHeader) ValidateBasic() error {
	return h.Header().ValidateBasic()
}

func (h *LightHeader) Header() *tmtypes.Header {
	return &tmtypes.Header{
		Version:            tmversion.Consensus(*h.Version),
		ChainID:            h.ChainId,
		Height:             h.Height,
		Time:               h.Time.Time(),
		LastBlockID:        *h.LastBlockId.BlockID(),
		LastCommitHash:     h.LastCommitHash,
		DataHash:           h.DataHash,
		ValidatorsHash:     h.ValidatorsHash,
		NextValidatorsHash: h.NextValidatorsHash,
		ConsensusHash:      h.ConsensusHash,
		AppHash:            h.AppHash,
		LastResultsHash:    h.LastResultsHash,
		EvidenceHash:       h.EvidenceHash,
		ProposerAddress:    h.ProposerAddress,
	}
}
