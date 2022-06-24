package module

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/datachainlab/harmony-cosmos-bridge-demo/relayer/chains/harmony"
	"github.com/datachainlab/harmony-cosmos-bridge-demo/relayer/chains/harmony/cmd"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/spf13/cobra"
)

type Module struct{}

var _ config.ModuleI = (*Module)(nil)

// Name returns the name of the module
func (Module) Name() string {
	return "harmony"
}

// RegisterInterfaces register the module interfaces to protobuf Any.
func (Module) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	harmony.RegisterInterfaces(registry)
}

// GetCmd returns the command
func (Module) GetCmd(ctx *config.Context) *cobra.Command {
	return cmd.HarmonyCmd(ctx.Codec, ctx)
}
