package harmony

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	hmylctypes "github.com/datachainlab/ibc-harmony-client/modules/light-clients/harmony/types"
	"github.com/hyperledger-labs/yui-relayer/core"
)

// RegisterInterfaces register the module interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	hmylctypes.RegisterInterfaces(registry)

	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
	registry.RegisterImplementations(
		(*exported.Header)(nil),
		&types.Header{},
	)
}
