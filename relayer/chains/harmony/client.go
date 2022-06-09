package harmony

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	sdkrpc "github.com/harmony-one/go-sdk/pkg/rpc"
	v1 "github.com/harmony-one/go-sdk/pkg/rpc/v1"
	v2 "github.com/harmony-one/harmony/rpc/v2"
)

const (
	MethodGetFullHeader  = "hmyv2_getFullHeader"
	MethodEpochLastBlock = "hmyv2_epochLastBlock"
	MethodGetEpoch       = "hmyv2_getEpoch"
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
	invalidRes := uint64(0)

	val, err := c.sendRPC(v1.Method.BlockNumber, nil)
	if err != nil {
		return invalidRes, err
	}
	bns, ok := val.(string)
	if !ok {
		return invalidRes, errors.New("could not get the latest block number")
	}
	return hexutil.DecodeUint64(bns)
}

func (c *Client) FullHeader(ctx context.Context, height uint64) (*v2.BlockHeader, error) {
	var heightArg string
	if height >= 0 {
		heightArg = strconv.FormatUint(height, 10)
	} else {
		heightArg = "latest"
	}
	val, err := c.sendRPC(MethodGetFullHeader, []interface{}{heightArg})
	if err != nil {
		return nil, err
	}

	jsonStr, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var header v2.BlockHeader
	if err := json.Unmarshal(jsonStr, &header); err != nil {
		return nil, err
	}
	return &header, nil
}

// Last block number of given epoch
func (c *Client) EpochLastBlock(ctx context.Context, epoch *big.Int) (*big.Int, error) {
	if epoch == nil {
		return nil, errors.New("epoch is null")
	}
	val, err := c.sendRPC(MethodEpochLastBlock, []interface{}{epoch.Uint64()})
	if err != nil {
		return nil, err
	}
	num, ok := val.(float64)
	if !ok {
		return nil, errors.New("could not get the last block of epoch")
	}
	bn, _ := big.NewFloat(num).Int(nil)
	return bn, nil
}

func (c *Client) GetLatestEpoch(ctx context.Context) (*big.Int, error) {
	val, err := c.sendRPC(MethodGetEpoch, []interface{}{})
	if err != nil {
		return nil, err
	}
	str, ok := val.(string)
	if !ok {
		return nil, errors.New("could not get the latest epoch")
	}
	num := new(big.Int)
	if _, ok := num.SetString(str, 10); !ok {
		return nil, errors.New("could not convert latest epoch to number")
	}
	return num, nil
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

func (c *Client) sendRPC(meth string, params []interface{}) (interface{}, error) {
	rep, err := c.messenger.SendRPC(meth, params)
	if err != nil {
		return nil, err
	}
	val, ok := rep["result"]
	if !ok {
		return nil, errors.New("invalid rpc response")
	}
	return val, nil
}
