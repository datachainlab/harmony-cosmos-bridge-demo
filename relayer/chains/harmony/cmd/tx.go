package cmd

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	"github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/datachainlab/harmony-cosmos-bridge-demo/relayer/chains/harmony"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/hyperledger-labs/yui-relayer/core"
	"github.com/spf13/cobra"
)

// txCmd represents the chain command
func txCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Tx Commands",
		Long:  "Commands to send tx.",
	}

	cmd.AddCommand(
		depositCmd(ctx),
		xfersend(ctx),
	)
	return cmd
}

// depositCmd represents the chain command
func depositCmd(ctx *config.Context) *cobra.Command {
	c := &cobra.Command{
		Use:   "deposit [chain-id]",
		Short: "Depsit Commands",
		Long:  "Commands to deposit token to ICS20TransferBank",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := ctx.Config.GetChain(args[0])
			if err != nil {
				return err
			}
			chain, ok := c.ChainI.(*harmony.Chain)
			if !ok {
				return errors.New("invalid chain-id")
			}
			owner, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			bankId, err := cmd.Flags().GetString(flagBankId)
			if err != nil {
				return err
			}
			balance, err := chain.QueryBankBalance(common.HexToAddress(owner), bankId)
			if err != nil {
				return err
			}
			fmt.Printf("%d\n", balance)
			return nil

		},
	}
	return c
}

func xfersend(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [path-name] [chain-id]",
		Short: "Initiate a transfer from one chain to another",
		Long: "Sends the first step to transfer tokens in an IBC transfer." +
			" The created packet must be relayed to another chain",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := ctx.Config.Paths.Get(args[0])
			if err != nil {
				return err
			}
			c, err := ctx.Config.GetChain(args[1])
			if err != nil {
				return err
			}
			_, ok := c.ChainI.(*harmony.Chain)
			if !ok {
				return errors.New("invalid chain-id")
			}

			amount, err := cmd.Flags().GetUint64(flagAmount)
			if err != nil {
				return err
			}
			// XXX want to support all denom format
			d, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}
			denom := transfertypes.ParseDenomTrace(d)
			token := sdk.Coin{
				Denom:  denom.GetFullDenomPath(),
				Amount: sdk.Int(sdk.NewUint(amount)),
			}
			if denom.Path != "" {
				token.Denom = denom.IBCDenom()
			}

			// Bech32 address string
			receiver, err := cmd.Flags().GetString(flagReceiver)
			if err != nil {
				return err
			}
			receiverAcc, err := sdk.AccAddressFromBech32(receiver)
			if err != nil {
				return err
			}
			fmt.Printf("receiver: %s", hexutil.Encode(receiverAcc.Bytes()))

			tx := core.RelayMsgs{
				Src: []sdk.Msg{},
				Dst: []sdk.Msg{
					transfertypes.NewMsgTransfer(
						path.Dst.PortID,
						path.Dst.ChannelID,
						token,
						"", // not used
						hexutil.Encode(receiverAcc.Bytes()),
						// TODO timeout height
						types.NewHeight(0, 10000),
						// TODO timeout timestamp
						uint64(0),
					),
				},
			}
			if tx.Send(nil, c); !tx.Succeeded {
				return fmt.Errorf("failed to send transfer message")
			}
			return nil
		},
	}
	return sendTransferFlags(cmd)
}
