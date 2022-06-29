package harmony

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	sdkcommon "github.com/harmony-one/go-sdk/pkg/common"
	"github.com/harmony-one/harmony/numeric"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.ChainConfigI = (*ChainConfig)(nil)
var _ core.ProverConfigI = (*ProverConfig)(nil)

func (c ChainConfig) Build() (core.ChainI, error) {
	return NewChain(c)
}

func (c ChainConfig) IBCHostAddress() common.Address {
	return common.HexToAddress(c.IbcHostAddress)
}

func (c ChainConfig) IBCHandlerAddress() common.Address {
	return common.HexToAddress(c.IbcHandlerAddress)
}

func (c ChainConfig) SimpleTokenAddress() common.Address {
	return common.HexToAddress(c.TokenAddress)
}

func (c ChainConfig) ICS20BankAddress() common.Address {
	return common.HexToAddress(c.Ics20BankAddress)
}

func (c ChainConfig) ChainID() (*sdkcommon.ChainID, error) {
	return sdkcommon.StringToChainID(c.HarmonyChainId)
}

func (c ChainConfig) GasPriceDec() numeric.Dec {
	return numeric.NewDec(c.GasPrice)
}

func (c ProverConfig) Build(chain core.ChainI) (core.ProverI, error) {
	hmyChain, ok := chain.(*Chain)
	if !ok {
		return nil, fmt.Errorf("invalid chain type")
	}
	return NewProver(hmyChain, c)
}

func (c ProverConfig) TrustingPeriodDuration() (time.Duration, error) {
	return time.ParseDuration(c.TrustingPeriod)
}
