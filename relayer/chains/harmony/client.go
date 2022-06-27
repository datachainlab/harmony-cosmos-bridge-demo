package harmony

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// FullHeader returns the harmony full header for the given height.
// The complete header can be used to calculate the hash value.
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

// EpochLastBlockNumber returns the last block number of the given epoch.
// Note that it also returns the block number for a future epoch.
func (c *Client) EpochLastBlockNumber(ctx context.Context, epoch uint64) (uint64, error) {
	val, err := c.sendRPC(MethodEpochLastBlock, []interface{}{epoch})
	if err != nil {
		return 0, err
	}
	// TODO DEBUG remove
	fmt.Printf("EpochLastBlock type: %T", val)
	// TODO modify
	num, ok := val.(float64)
	if !ok {
		return 0, errors.New("could not get the last block of epoch")
	}
	// TODO modify
	bn, _ := big.NewFloat(num).Int(nil)
	return bn.Uint64(), nil
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
		return nil, fmt.Errorf("rpc %s with params %v failed: %w", meth, params, err)
	}
	val, ok := rep["result"]
	if !ok {
		return nil, fmt.Errorf("rpc %s with params %v returns invalid response", meth, params)
	}
	return val, nil
}
