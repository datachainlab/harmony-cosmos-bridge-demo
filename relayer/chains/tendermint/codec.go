package tendermint

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/datachainlab/harmony-cosmos-bridge-demo/relayer/chains/tendermint/types"
	"github.com/hyperledger-labs/yui-relayer/core"
)

// RegisterInterfaces register the module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
	types.RegisterInterfaces(registry)
}
