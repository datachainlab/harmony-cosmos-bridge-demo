package cmd_test

import (
	"fmt"
	"testing"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/datachainlab/harmony-cosmos-bridge-demo/tests/chains/tendermint/simapp"
	"github.com/datachainlab/harmony-cosmos-bridge-demo/tests/chains/tendermint/simapp/simd/cmd"
	"github.com/stretchr/testify/require"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",        // Test the init cmd
		"simapp-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
	})

	require.NoError(t, svrcmd.Execute(rootCmd, simapp.DefaultNodeHome))
}
