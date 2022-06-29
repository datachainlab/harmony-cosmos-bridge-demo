package cmd

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagHash                = "hash"
	flagForce               = "force"
	flagTimeoutHeightOffset = "timeout-height-offset"
	flagTimeoutTimeOffset   = "timeout-time-offset"
)

func lightFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Int64(flags.FlagHeight, -1, "Trusted header's height")
	cmd.Flags().BytesHexP(flagHash, "x", []byte{}, "Trusted header's hash")
	if err := viper.BindPFlag(flags.FlagHeight, cmd.Flags().Lookup(flags.FlagHeight)); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(flagHash, cmd.Flags().Lookup(flagHash)); err != nil {
		panic(err)
	}
	return cmd
}

func forceFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP(flagForce, "f", false, "option to force non-standard behavior such as initialization of light client from configured chain or generation of new path") //nolint:lll
	if err := viper.BindPFlag(flagForce, cmd.Flags().Lookup(flagForce)); err != nil {
		panic(err)
	}
	return cmd
}

func timeoutFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Uint64P(flagTimeoutHeightOffset, "y", 0, "set timeout height offset for ")
	cmd.Flags().DurationP(flagTimeoutTimeOffset, "c", time.Duration(0), "specify the path to relay over")
	if err := viper.BindPFlag(flagTimeoutHeightOffset, cmd.Flags().Lookup(flagTimeoutHeightOffset)); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(flagTimeoutTimeOffset, cmd.Flags().Lookup(flagTimeoutTimeOffset)); err != nil {
		panic(err)
	}
	return cmd
}
