package cmd

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/spf13/cobra"
)

func HarmonyCmd(m codec.Codec, ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "harmony",
		Short: "manage harmony configurations",
	}

	cmd.AddCommand(
		queryCmd(ctx),
		txCmd(ctx),
	)

	return cmd
}
