harmony-cosmos-bridge-demo
---

# Prerequisites

## Env

```
export LOCAL_SHARD_0_URL=http://localhost:9599
export LOCAL_SHARD_1_URL=http://localhost:9598
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
