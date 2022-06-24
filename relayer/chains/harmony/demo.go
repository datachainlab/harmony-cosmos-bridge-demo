package harmony

import (
	"context"
	"log"
	"math/big"
	"strings"

	transfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/harmony-one/go-sdk/pkg/transaction"
	harmonytypes "github.com/harmony-one/harmony/core/types"
	"github.com/harmony-one/harmony/numeric"
)

const (
	msgTxMsgTransfer = "sendTransfer"
)

func (c *Chain) QueryTokenBalance(address common.Address) (*big.Int, error) {
	return c.simpleToken.BalanceOf(c.CallOpts(context.Background(), -1), address)
}

func (c *Chain) QueryBankBalance(address common.Address, id string) (*big.Int, error) {
	idLower := strings.ToLower(id)
	return c.ics20Bank.BalanceOf(c.CallOpts(context.Background(), -1), address, idLower)
}

func (c *Chain) TxMsgTransfer(msg *transfertypes.MsgTransfer) (*harmonytypes.Transaction, error) {
	denomLower := strings.ToLower(msg.Token.Denom)
	return c.txIcs20TransferBank(
		msgTxMsgTransfer,
		denomLower,
		msg.Token.Amount.Uint64(),
		common.HexToAddress(msg.Receiver),
		msg.SourcePort,
		msg.SourceChannel,
		msg.TimeoutHeight.RevisionHeight,
	)
}

func (c *Chain) txIcs20TransferBank(method string, params ...interface{}) (*harmonytypes.Transaction, error) {
	input, err := c.ics20TransferBankAbi.Pack(method, params...)
	if err != nil {
		log.Println("abi.Pack error")
		return nil, err
	}
	account, err := c.getAccount()
	if err != nil {
		return nil, err
	}
	if err = c.keyStore.Unlock(account, ""); err != nil {
		return nil, err
	}
	controller := transaction.NewController(c.client.messenger, c.keyStore, &account, *c.chainId)
	nonce := transaction.GetNextPendingNonce(account.Address.Hex(), c.client.messenger)
	err = controller.ExecuteTransaction(nonce, c.config.GasLimit, &c.config.Ics20TransferBankAddress, c.config.ShardId, c.config.ShardId, numeric.NewDec(0), c.config.GasPriceDec(), input)
	if err != nil {
		log.Println("config.GasLimit", c.config.GasLimit)
		return nil, err
	}
	if err = c.keyStore.Lock(account.Address); err != nil {
		panic(err)
	}
	return controller.TransactionInfo(), nil
}
