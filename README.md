harmony-cosmos-bridge-demo
---

This is a demonstration of token transfer using IBC between a Harmony network and a Cosmos-based blockchain.

# IBC Light Client

- We use [yui-ibc-solidity](https://github.com/hyperledger-labs/yui-ibc-solidity) for IBC core and our customed [tendermint-sol](https://github.com/datachainlab/tendermint-sol/tree/use-ibc-sol-hmy) as a Tendermint light client on Harmony
- We use [ibc-harmony-client](https://github.com/datachainlab/ibc-harmony-client) as a Harmony light client on Cosmos-based blockchain


# Directory structure
- contracts ... contracts and migration scripts for harmony
- relayer ... IBC relayer using [yui-relayer](https://github.com/hyperledger-labs/yui-relayer)
    - chains/harmony ... harmony modules
    - chains/tendermint ... tendermint modules for tendermint-sol
- tests ... test environments
    - cases ... E2E test case
    - chains/harmony ... configs for harmony localnet
    - chains/tendermint ... configs for tendermint localnet


# Setup

## Prerequisites

- bash 4.0 or higher
- golang 1.16 or higher
- node.js 16

## Env

Example:

```
# assume that RPC port offset is 500
export HARMONY_LOCAL_SHARD_0_URL=http://localhost:9598
export HARMONY_LOCAL_SHARD_1_URL=http://localhost:9596

export HARMONY_GAS_LIMIT=100000000
export HARMONY_GAS_PRICE=1000000000

# used for deploying contracts
# cf. https://github.com/harmony-one/harmony/pull/3332
export HARMONY_LOCAL_PRIVATE_KEY: '0x1f84c95ac16e6a50f08d44c7bde7aff8742212fda6e4321fde48bf83bef266dc'
```

Harmony localnet url is configured based on [tests/chains/harmony/configs/localnet_deploy.config](tests/chains/harmony/docker/configs/localnet_deploy.config).


## Building Relayer

Relayer needs libraries for harmony.
At first, we need to build the libraries.

```
make clone-harmony
make build-harmony
```

Then, we will get a relayer with the following command.

```
make build-relayer
```

## Preparing Harmony Local Network

The following commands creates a Harmony localnet docker image with deployed contracts.

```
cd tests/chains/harmony
make docker-image
```

## Preparing Cosmos Local Network

The following command creates a Cosmos local network image.

```
cd tests/chains/tendermint
make docker-image
```

# E2E

The following commands brings up the two networks, performs an IBC Handshake, and transfers token.

```
cd tests/cases/tm2harmony
make network
make test
make network-down
```
