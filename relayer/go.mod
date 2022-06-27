module github.com/datachainlab/harmony-cosmos-bridge-demo/relayer

go 1.16

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/ibc-harmony-client v0.0.0-20220627020344-c0def938d048
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/harmony-one/go-sdk v1.2.8
	github.com/harmony-one/harmony v1.10.3-0.20220129011036-4ea9072e5eda
	// use go bindings with geth v1.9.10, which works with solidity 6.0+
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20220624103310-247f169b23ce
	github.com/hyperledger-labs/yui-relayer v0.1.1-0.20211209032245-495b5eed40e2
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/valyala/fasthttp v1.4.0 // indirect
)

replace (
	// for harmony-one/harmony
	github.com/ethereum/go-ethereum v1.9.25 => github.com/ethereum/go-ethereum v1.9.9
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/harmony-one/go-sdk => github.com/datachainlab/go-sdk v1.2.9-0.20220106070458-8ce5f5c807b2
)
