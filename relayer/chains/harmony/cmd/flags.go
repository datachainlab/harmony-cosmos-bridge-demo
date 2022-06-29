package cmd

import (
	"github.com/spf13/cobra"
)

const (
	flagOwner               = "owner"
	flagSigner              = "signer"
	flagReceiver            = "receiver"
	flagAmount              = "amount"
	flagDenom               = "denom"
	flagBankId              = "bank-id"
	flagTimeoutHeightOffset = "timeout-height-offset"
	flagTimeoutTimeOffset   = "timeout-time-offset"
)

func ownerFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().String(flagOwner, "", "owner hex address string")
	_ = cmd.MarkFlagRequired(flagOwner)
	return cmd
}

func sendTransferFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().String(flagDenom, "", "denom")
	cmd.Flags().String(flagReceiver, "", "receiver address")
	cmd.Flags().Uint64(flagAmount, 0, "amount")
	_ = cmd.MarkFlagRequired(flagDenom)
	_ = cmd.MarkFlagRequired(flagReceiver)
	_ = cmd.MarkFlagRequired(flagAmount)
	return cmd
}

func bankIdFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().String(flagBankId, "", "bank id")
	_ = cmd.MarkFlagRequired(flagBankId)
	return cmd
}
