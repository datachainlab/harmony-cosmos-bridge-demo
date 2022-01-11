harmony-cosmos-bridge-demo
---

# Prerequisites

## Env

```
export HARMONY_LOCAL_SHARD_0_URL=http://localhost:9599
export HARMONY_LOCAL_SHARD_1_URL=http://localhost:9598
export HARMONY_GAS_LIMIT=100000000
export HARMONY_GAS_PRICE=1000000000
# cf. https://github.com/harmony-one/harmony/pull/3332
export HARMONY_LOCAL_PRIVATE_KEY: '0x1f84c95ac16e6a50f08d44c7bde7aff8742212fda6e4321fde48bf83bef266dc'
```

## Harmony Localnet

Clone the necessary repositories and set up a harmony localnet

```
make setup-harmony
```

# Setup

## Relayer

TODO

## Harmony Localnet

Compile and deploy contracts:

```
make compile-contracts
make debug-harmony &
make deploy-contracts-shard0 # or make deploy-contracts-shard1
```

# E2E

TODO
