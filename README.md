harmony-cosmos-bridge-demo
---

# Prerequisites

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

Harmony localnet url is configured based on [tests/chains/harmony/configs/localnet_deploy.config](tests/chains/harmony/configs/localnet_deploy.config).

# Setup

## Relayer

Relayer needs libraries for harmony.
Clone the necessary repositories and build them.

```
make clone-harmony
make build-harmony
```

```
make build-relayer
```

## Harmony Localnet

### Building image

```
cd tests/chains/harmony
make docker-image
```

### Starting localnet

```
cd tests/cases/tm2harmony
make network-harmony
```

### Deploying contracts

```
make deploy-contracts-shard0
```

# E2E

Currently up to IBC Handshake with mock client.

```
cd tests/cases/tm2harmony
make network
make test
make network-down
```
