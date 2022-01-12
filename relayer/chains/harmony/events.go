package harmony

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	v1 "github.com/harmony-one/go-sdk/pkg/rpc/v1"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
)

var (
	abiSendPacket,
	abiWriteAcknowledgement,
	abiGeneratedClientIdentifier,
	abiGeneratedConnectionIdentifier,
	abiGeneratedChannelIdentifier abi.Event
)

func init() {
	parsedHandlerABI, err := abi.JSON(strings.NewReader(ibchandler.IbchandlerABI))
	if err != nil {
		panic(err)
	}
	parsedHostABI, err := abi.JSON(strings.NewReader(ibchost.IbchostABI))
	if err != nil {
		panic(err)
	}
	abiSendPacket = parsedHandlerABI.Events["SendPacket"]
	abiWriteAcknowledgement = parsedHandlerABI.Events["WriteAcknowledgement"]
	abiGeneratedClientIdentifier = parsedHostABI.Events["GeneratedClientIdentifier"]
	abiGeneratedConnectionIdentifier = parsedHostABI.Events["GeneratedConnectionIdentifier"]
	abiGeneratedChannelIdentifier = parsedHostABI.Events["GeneratedChannelIdentifier"]
}

func (chain *Chain) findPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
	sequence uint64,
) (*chantypes.Packet, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			chain.config.IBCHandlerAddress(),
		},
		Topics: [][]common.Hash{{
			abiSendPacket.ID,
		}},
	}
	logsData, err := chain.findLogsData(ctx, query)
	if err != nil {
		return nil, err
	}
	for _, data := range logsData {
		if values, err := abiSendPacket.Inputs.Unpack(data); err != nil {
			return nil, err
		} else {
			p := values[0].(struct {
				Sequence           uint64  "json:\"sequence\""
				SourcePort         string  "json:\"source_port\""
				SourceChannel      string  "json:\"source_channel\""
				DestinationPort    string  "json:\"destination_port\""
				DestinationChannel string  "json:\"destination_channel\""
				Data               []uint8 "json:\"data\""
				TimeoutHeight      struct {
					RevisionNumber uint64 "json:\"revision_number\""
					RevisionHeight uint64 "json:\"revision_height\""
				} "json:\"timeout_height\""
				TimeoutTimestamp uint64 "json:\"timeout_timestamp\""
			})
			if p.SourcePort == sourcePortID && p.SourceChannel == sourceChannel && p.Sequence == sequence {
				return &chantypes.Packet{
					Sequence:           p.Sequence,
					SourcePort:         p.SourcePort,
					SourceChannel:      p.SourceChannel,
					DestinationPort:    p.DestinationPort,
					DestinationChannel: p.DestinationChannel,
					Data:               p.Data,
					TimeoutHeight:      clienttypes.Height(p.TimeoutHeight),
					TimeoutTimestamp:   p.TimeoutTimestamp,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("packet not found: sourcePortID=%v sourceChannel=%v sequence=%v", sourcePortID, sourceChannel, sequence)
}

// getAllPackets returns all packets from events
func (chain *Chain) getAllPackets(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
) ([]*chantypes.Packet, error) {
	var packets []*chantypes.Packet

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			chain.config.IBCHandlerAddress(),
		},
		Topics: [][]common.Hash{{
			abiSendPacket.ID,
		}},
	}
	logsData, err := chain.findLogsData(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, data := range logsData {
		if values, err := abiSendPacket.Inputs.Unpack(data); err != nil {
			return nil, err
		} else {
			p := values[0].(struct {
				Sequence           uint64  "json:\"sequence\""
				SourcePort         string  "json:\"source_port\""
				SourceChannel      string  "json:\"source_channel\""
				DestinationPort    string  "json:\"destination_port\""
				DestinationChannel string  "json:\"destination_channel\""
				Data               []uint8 "json:\"data\""
				TimeoutHeight      struct {
					RevisionNumber uint64 "json:\"revision_number\""
					RevisionHeight uint64 "json:\"revision_height\""
				} "json:\"timeout_height\""
				TimeoutTimestamp uint64 "json:\"timeout_timestamp\""
			})
			if p.SourcePort == sourcePortID && p.SourceChannel == sourceChannel {
				packet := &chantypes.Packet{
					Sequence:           p.Sequence,
					SourcePort:         p.SourcePort,
					SourceChannel:      p.SourceChannel,
					DestinationPort:    p.DestinationPort,
					DestinationChannel: p.DestinationChannel,
					Data:               p.Data,
					TimeoutHeight:      clienttypes.Height(p.TimeoutHeight),
					TimeoutTimestamp:   p.TimeoutTimestamp,
				}
				packets = append(packets, packet)
			}
		}
	}
	return packets, nil
}

func (chain *Chain) findAcknowledgement(
	ctx context.Context,
	dstPortID string,
	dstChannel string,
	sequence uint64,
) ([]byte, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			chain.config.IBCHandlerAddress(),
		},
		Topics: [][]common.Hash{{
			abiWriteAcknowledgement.ID,
		}},
	}
	logsData, err := chain.findLogsData(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, data := range logsData {
		if values, err := abiWriteAcknowledgement.Inputs.Unpack(data); err != nil {
			return nil, err
		} else {
			if len(values) != 4 {
				return nil, fmt.Errorf("unexpected values: %v", values)
			}
			if dstPortID == values[0].(string) && dstChannel == values[1].(string) && sequence == values[2].(uint64) {
				return values[3].([]byte), nil
			}
		}
	}

	return nil, fmt.Errorf("ack not found: dstPortID=%v dstChannel=%v sequence=%v", dstPortID, dstChannel, sequence)
}

type PacketAcknowledgement struct {
	Sequence uint64
	Data     []byte
}

func (chain *Chain) getAllAcknowledgements(
	ctx context.Context,
	dstPortID string,
	dstChannel string,
) ([]PacketAcknowledgement, error) {
	var acks []PacketAcknowledgement
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			chain.config.IBCHandlerAddress(),
		},
		Topics: [][]common.Hash{{
			abiWriteAcknowledgement.ID,
		}},
	}
	logsData, err := chain.findLogsData(ctx, query)
	if err != nil {
		return nil, err
	}
	for _, data := range logsData {
		if values, err := abiWriteAcknowledgement.Inputs.Unpack(data); err != nil {
			return nil, err
		} else {
			if len(values) != 4 {
				return nil, fmt.Errorf("unexpected values: %v", values)
			}
			if dstPortID == values[0].(string) && dstChannel == values[1].(string) {
				acks = append(acks, PacketAcknowledgement{
					Sequence: values[2].(uint64),
					Data:     values[3].([]byte),
				})
			}
		}
	}
	return acks, nil
}

func (chain *Chain) findLogsData(ctx context.Context, q ethereum.FilterQuery) ([][]byte, error) {
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	rep, err := chain.client.messenger.SendRPC(v1.Method.GetPastLogs, []interface{}{arg})
	if err != nil {
		return nil, err
	}
	result, ok := rep["result"]
	if !ok {
		return nil, errors.New("invalid request")
	}

	rs, ok := result.([]interface{})
	if !ok {
		return nil, errors.New("can't convert result to slice")
	}
	// need all data?
	data := make([][]byte, len(rs))
	for i, r := range rs {
		resmap, ok := r.(map[string]interface{})
		if !ok {
			return nil, errors.New("can't convert to map")
		}
		dataStr, ok := resmap["data"].(string)
		if !ok {
			return nil, errors.New("can't convert data")
		}
		bz, err := hexutil.Decode(dataStr)
		if err != nil {
			return nil, err
		}
		data[i] = bz
	}
	return data, nil
}
