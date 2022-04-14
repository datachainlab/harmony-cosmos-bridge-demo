package main

import (
	"log"

	harmony "github.com/datachainlab/harmony-cosmos-bridge-demo/relayer/chains/harmony/module"
	tendermint "github.com/datachainlab/harmony-cosmos-bridge-demo/relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
	mock "github.com/hyperledger-labs/yui-relayer/provers/mock/module"
)

func main() {
	if err := cmd.Execute(
		harmony.Module{},
		tendermint.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
