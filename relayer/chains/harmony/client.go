package harmony

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	sdkrpc "github.com/harmony-one/go-sdk/pkg/rpc"
	v1 "github.com/harmony-one/go-sdk/pkg/rpc/v1"
)

type Client struct {
	messenger *sdkrpc.HTTPMessenger
}

func NewHarmonyClient(endpoint string) *Client {
	messenger := sdkrpc.NewHTTPHandler(endpoint)
	return &Client{
		messenger: messenger,
	}
}

func NewETHClient(endpoint string) (*ethclient.Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(conn), nil
}

// BlockNumber returns the most recent block number
func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	res := uint64(0)

	rep, err := c.messenger.SendRPC(v1.Method.BlockNumber, nil)
	if err != nil {
		return res, err
	}
	val, ok := rep["result"]
	if !ok {
		return res, errors.New("invalid response")
	}
	bns, ok := val.(string)
	if !ok {
		return res, errors.New("invalid result")
	}
	return hexutil.DecodeUint64(bns)
}

func (chain *Chain) CallOpts(ctx context.Context, height int64) *bind.CallOpts {
	account, err := chain.getAccount()
	if err != nil {
		return &bind.CallOpts{
			Context: ctx,
		}
	}
	opts := &bind.CallOpts{
		From:    account.Address,
		Context: ctx,
	}
	if height > 0 {
		opts.BlockNumber = big.NewInt(height)
	}
	return opts
}
