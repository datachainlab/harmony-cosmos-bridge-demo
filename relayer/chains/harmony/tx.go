package harmony

import (
	"errors"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	proto "github.com/gogo/protobuf/proto"
	"github.com/harmony-one/go-sdk/pkg/transaction"
	"github.com/harmony-one/harmony/accounts"
	"github.com/harmony-one/harmony/accounts/abi"
	harmonytypes "github.com/harmony-one/harmony/core/types"
	"github.com/harmony-one/harmony/numeric"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

const (
	methodCreateClient          = "createClient"
	methodUpdateClient          = "updateClient"
	methodConnectionOpenInit    = "connectionOpenInit"
	methodConnectionOpenTry     = "connectionOpenTry"
	methodConnectionOpenAck     = "connectionOpenAck"
	methodConnectionOpenConfirm = "connectionOpenConfirm"
	methodChannelOpenInit       = "channelOpenInit"
	methodChannelOpenTry        = "channelOpenTry"
	methodChannelOpenAck        = "channelOpenAck"
	methodChannelOpenConfirm    = "channelOpenConfirm"
	methodRecvPacket            = "recvPacket"
	methodAcknowledgement       = "acknowledgePacket"
)

// SendMsgs sends msgs to the chain
func (c *Chain) SendMsgs(msgs []sdk.Msg) ([]byte, error) {
	for _, msg := range msgs {
		var (
			err error
		)
		switch msg := msg.(type) {
		case *clienttypes.MsgCreateClient:
			_, err = c.TxCreateClient(msg)
		case *clienttypes.MsgUpdateClient:
			_, err = c.TxUpdateClient(msg)
		case *conntypes.MsgConnectionOpenInit:
			_, err = c.TxConnectionOpenInit(msg)
		case *conntypes.MsgConnectionOpenTry:
			_, err = c.TxConnectionOpenTry(msg)
		case *conntypes.MsgConnectionOpenAck:
			_, err = c.TxConnectionOpenAck(msg)
		case *conntypes.MsgConnectionOpenConfirm:
			_, err = c.TxConnectionOpenConfirm(msg)
		case *chantypes.MsgChannelOpenInit:
			_, err = c.TxChannelOpenInit(msg)
		case *chantypes.MsgChannelOpenTry:
			_, err = c.TxChannelOpenTry(msg)
		case *chantypes.MsgChannelOpenAck:
			_, err = c.TxChannelOpenAck(msg)
		case *chantypes.MsgChannelOpenConfirm:
			_, err = c.TxChannelOpenConfirm(msg)
		case *chantypes.MsgRecvPacket:
			_, err = c.TxRecvPacket(msg)
		case *chantypes.MsgAcknowledgement:
			_, err = c.TxAcknowledgement(msg)
		default:
			panic("illegal msg type")
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// Send sends msgs to the chain and logging a result of it
// It returns a boolean value whether the result is success
func (c *Chain) Send(msgs []sdk.Msg) bool {
	_, err := c.SendMsgs(msgs)
	if err != nil {
		log.Println("harmony: failed to send:", err)
	}
	return err == nil
}

func (c *Chain) TxCreateClient(msg *clienttypes.MsgCreateClient) (*harmonytypes.Transaction, error) {
	var clientState exported.ClientState
	if err := c.codec.UnpackAny(msg.ClientState, &clientState); err != nil {
		return nil, err
	}
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	consensusStateBytes, err := proto.Marshal(msg.ConsensusState)
	if err != nil {
		return nil, err
	}
	log.Printf("TxCreateClient: ClientType: %+v, Height: %+v\n", clientState.ClientType(),
		clientState.GetLatestHeight().GetRevisionHeight())
	return c.txIbcHandler(methodCreateClient, ibchandler.IBCMsgsMsgCreateClient{
		ClientType:          clientState.ClientType(),
		Height:              clientState.GetLatestHeight().GetRevisionHeight(),
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	})
}

func (c *Chain) TxUpdateClient(msg *clienttypes.MsgUpdateClient) (*harmonytypes.Transaction, error) {
	headerBytes, err := proto.Marshal(msg.Header)
	if err != nil {
		return nil, err
	}
	return c.txIbcHandler(methodUpdateClient, ibchandler.IBCMsgsMsgUpdateClient{
		ClientId: msg.ClientId,
		Header:   headerBytes,
	})
}

func (c *Chain) TxConnectionOpenInit(msg *conntypes.MsgConnectionOpenInit) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodConnectionOpenInit, ibchandler.IBCMsgsMsgConnectionOpenInit{
		ClientId: msg.ClientId,
		Counterparty: ibchandler.CounterpartyData{
			ClientId:     msg.Counterparty.ClientId,
			ConnectionId: msg.Counterparty.ConnectionId,
			Prefix:       ibchandler.MerklePrefixData(msg.Counterparty.Prefix),
		},
		DelayPeriod: msg.DelayPeriod,
	})
}

func (c *Chain) TxConnectionOpenTry(msg *conntypes.MsgConnectionOpenTry) (*harmonytypes.Transaction, error) {
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	var versions []ibchandler.VersionData
	for _, v := range msg.CounterpartyVersions {
		versions = append(versions, ibchandler.VersionData(*v))
	}
	return c.txIbcHandler(methodConnectionOpenTry, ibchandler.IBCMsgsMsgConnectionOpenTry{
		PreviousConnectionId: msg.PreviousConnectionId,
		Counterparty: ibchandler.CounterpartyData{
			ClientId:     msg.Counterparty.ClientId,
			ConnectionId: msg.Counterparty.ConnectionId,
			Prefix:       ibchandler.MerklePrefixData(msg.Counterparty.Prefix),
		},
		DelayPeriod:          msg.DelayPeriod,
		ClientId:             msg.ClientId,
		ClientStateBytes:     clientStateBytes,
		CounterpartyVersions: versions,
		ProofInit:            msg.ProofInit,
		ProofClient:          msg.ProofClient,
		ProofConsensus:       msg.ProofConsensus,
		ProofHeight:          msg.ProofHeight.RevisionHeight,
		ConsensusHeight:      msg.ConsensusHeight.RevisionHeight,
	})
}

func (c *Chain) TxConnectionOpenAck(msg *conntypes.MsgConnectionOpenAck) (*harmonytypes.Transaction, error) {
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	return c.txIbcHandler(methodConnectionOpenAck, ibchandler.IBCMsgsMsgConnectionOpenAck{
		ConnectionId:     msg.ConnectionId,
		ClientStateBytes: clientStateBytes,
		Version: ibchandler.VersionData{
			Identifier: msg.Version.Identifier,
			Features:   msg.Version.Features,
		},
		CounterpartyConnectionID: msg.CounterpartyConnectionId,
		ProofTry:                 msg.ProofTry,
		ProofClient:              msg.ProofClient,
		ProofConsensus:           msg.ProofConsensus,
		ProofHeight:              msg.ProofHeight.RevisionHeight,
		ConsensusHeight:          msg.ConsensusHeight.RevisionHeight,
	})
}

func (c *Chain) TxConnectionOpenConfirm(msg *conntypes.MsgConnectionOpenConfirm) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodConnectionOpenConfirm, ibchandler.IBCMsgsMsgConnectionOpenConfirm{
		ConnectionId: msg.ConnectionId,
		ProofAck:     msg.ProofAck,
		ProofHeight:  msg.ProofHeight.RevisionHeight,
	})
}

func (c *Chain) TxChannelOpenInit(msg *chantypes.MsgChannelOpenInit) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodChannelOpenInit, ibchandler.IBCMsgsMsgChannelOpenInit{
		PortId: msg.PortId,
		Channel: ibchandler.ChannelData{
			State:          uint8(msg.Channel.State),
			Ordering:       uint8(msg.Channel.Ordering),
			Counterparty:   ibchandler.ChannelCounterpartyData(msg.Channel.Counterparty),
			ConnectionHops: msg.Channel.ConnectionHops,
			Version:        msg.Channel.Version,
		},
	})
}

func (c *Chain) TxChannelOpenTry(msg *chantypes.MsgChannelOpenTry) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodChannelOpenTry, ibchandler.IBCMsgsMsgChannelOpenTry{
		PortId:            msg.PortId,
		PreviousChannelId: msg.PreviousChannelId,
		Channel: ibchandler.ChannelData{
			State:          uint8(msg.Channel.State),
			Ordering:       uint8(msg.Channel.Ordering),
			Counterparty:   ibchandler.ChannelCounterpartyData(msg.Channel.Counterparty),
			ConnectionHops: msg.Channel.ConnectionHops,
			Version:        msg.Channel.Version,
		},
		CounterpartyVersion: msg.CounterpartyVersion,
		ProofInit:           msg.ProofInit,
		ProofHeight:         msg.ProofHeight.RevisionHeight,
	})
}

func (c *Chain) TxChannelOpenAck(msg *chantypes.MsgChannelOpenAck) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodChannelOpenAck, ibchandler.IBCMsgsMsgChannelOpenAck{
		PortId:                msg.PortId,
		ChannelId:             msg.ChannelId,
		CounterpartyVersion:   msg.CounterpartyVersion,
		CounterpartyChannelId: msg.CounterpartyChannelId,
		ProofTry:              msg.ProofTry,
		ProofHeight:           msg.ProofHeight.RevisionHeight,
	})
}

func (c *Chain) TxChannelOpenConfirm(msg *chantypes.MsgChannelOpenConfirm) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodChannelOpenConfirm, ibchandler.IBCMsgsMsgChannelOpenConfirm{
		PortId:      msg.PortId,
		ChannelId:   msg.ChannelId,
		ProofAck:    msg.ProofAck,
		ProofHeight: msg.ProofHeight.RevisionHeight,
	})
}

func (c *Chain) TxRecvPacket(msg *chantypes.MsgRecvPacket) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodRecvPacket, ibchandler.IBCMsgsMsgPacketRecv{
		Packet: ibchandler.PacketData{
			Sequence:           msg.Packet.Sequence,
			SourcePort:         msg.Packet.SourcePort,
			SourceChannel:      msg.Packet.SourceChannel,
			DestinationPort:    msg.Packet.DestinationPort,
			DestinationChannel: msg.Packet.DestinationChannel,
			Data:               msg.Packet.Data,
			TimeoutHeight:      ibchandler.HeightData(msg.Packet.TimeoutHeight),
			TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		},
		Proof:       msg.ProofCommitment,
		ProofHeight: msg.ProofHeight.RevisionHeight,
	})
}

func (c *Chain) TxAcknowledgement(msg *chantypes.MsgAcknowledgement) (*harmonytypes.Transaction, error) {
	return c.txIbcHandler(methodAcknowledgement, ibchandler.IBCMsgsMsgPacketAcknowledgement{
		Packet: ibchandler.PacketData{
			Sequence:           msg.Packet.Sequence,
			SourcePort:         msg.Packet.SourcePort,
			SourceChannel:      msg.Packet.SourceChannel,
			DestinationPort:    msg.Packet.DestinationPort,
			DestinationChannel: msg.Packet.DestinationChannel,
			Data:               msg.Packet.Data,
			TimeoutHeight:      ibchandler.HeightData(msg.Packet.TimeoutHeight),
			TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		},
		Acknowledgement: msg.Acknowledgement,
		Proof:           msg.ProofAcked,
		ProofHeight:     msg.ProofHeight.RevisionHeight,
	})
}

func (c *Chain) txIbcHandler(method string, params ...interface{}) (*harmonytypes.Transaction, error) {
	//	return c.tx(c.config.IbcHandlerAddress, &c.ibcHandlerAbi, method, params)
	input, err := c.ibcHandlerAbi.Pack(method, params...)
	if err != nil {
		log.Println("abi.Pack error")
		return nil, err
	}
	account, err := c.getAccount()
	if err != nil {
		return nil, err
	}
	if err = c.keyStore.Unlock(account, ""); err != nil {
		return nil, err
	}
	controller := transaction.NewController(c.client.messenger, c.keyStore, &account, *c.chainId)
	// XXX or GetNextNonce
	nonce := transaction.GetNextPendingNonce(account.Address.Hex(), c.client.messenger)
	err = controller.ExecuteTransaction(nonce, c.config.GasLimit, &c.config.IbcHandlerAddress, c.config.ShardId, c.config.ShardId, numeric.NewDec(0), c.config.GasPriceDec(), input)
	if err != nil {
		log.Println("config.GasLimit", c.config.GasLimit)
		return nil, err
	}
	if err = c.keyStore.Lock(account.Address); err != nil {
		panic(err)
	}
	return controller.TransactionInfo(), nil

}

func (c *Chain) tx(to string, abi *abi.ABI, method string, params ...interface{}) (*harmonytypes.Transaction, error) {
	input, err := abi.Pack(method, params...)
	if err != nil {
		log.Println("abi.Pack error")
		return nil, err
	}
	account, err := c.getAccount()
	if err != nil {
		return nil, err
	}
	if err = c.keyStore.Unlock(account, ""); err != nil {
		return nil, err
	}
	controller := transaction.NewController(c.client.messenger, c.keyStore, &account, *c.chainId)
	// XXX or pending nonce
	nonce := transaction.GetNextNonce(account.Address.Hex(), c.client.messenger)
	err = controller.ExecuteTransaction(nonce, c.config.GasLimit, &to, c.config.ShardId, c.config.ShardId, numeric.NewDec(0), c.config.GasPriceDec(), input)
	if err != nil {
		return nil, err
	}
	if err = c.keyStore.Lock(account.Address); err != nil {
		panic(err)
	}
	return controller.TransactionInfo(), nil
}

func (chain *Chain) getAccount() (accounts.Account, error) {
	accs := chain.keyStore.Accounts()
	if len(accs) == 0 {
		return accounts.Account{}, errors.New("empty keystore")
	}
	return accs[0], nil
}
