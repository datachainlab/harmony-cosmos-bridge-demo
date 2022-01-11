package harmony

import (
	"context"
	"errors"

	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
)

func init() {
}

func (chain *Chain) findPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
	sequence uint64,
) (*chantypes.Packet, error) {
	return nil, errors.New("TODO: not implemented")
}

// getAllPackets returns all packets from events
func (chain *Chain) getAllPackets(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
) ([]*chantypes.Packet, error) {
	return []*chantypes.Packet{}, errors.New("TODO: not implemented")
}

func (chain *Chain) findAcknowledgement(
	ctx context.Context,
	dstPortID string,
	dstChannel string,
	sequence uint64,
) ([]byte, error) {
	return []byte{}, errors.New("TODO: not implemented")
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
	return []PacketAcknowledgement{}, errors.New("TODO: not implemented")
}
