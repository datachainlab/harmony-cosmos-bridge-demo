module github.com/datachainlab/harmony-cosmos-bridge-demo/tests/chains/tendermint

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/ibc-harmony-client v0.0.0-20220623084557-d600c9e6c9b0
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
)

replace (
	github.com/coinbase/rosetta-sdk-go => github.com/coinbase/rosetta-sdk-go v0.5.9
	github.com/ethereum/go-ethereum v1.9.25 => github.com/ethereum/go-ethereum v1.9.9
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
