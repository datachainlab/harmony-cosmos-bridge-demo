{
  "chain": {
    "@type": "/relayer.chains.harmony.config.ChainConfig",
    "chain_id": "ibc1",
    "harmony_chain_id": "testnet",
    "shard_id": 0,
    "shard_rpc_addr": "http://localhost:9598",
    "shard_private_key": "1f84c95ac16e6a50f08d44c7bde7aff8742212fda6e4321fde48bf83bef266dc",
    "beacon_rpc_addr": "http://localhost:9598",
    "beacon_private_key": "1f84c95ac16e6a50f08d44c7bde7aff8742212fda6e4321fde48bf83bef266dc",
    "ibc_host_address": "",
    "ibc_handler_address": "",
    "gas_limit": 5000000,
    "gas_price": 1000000000
  },
  "prover": {
    "@type": "/relayer.chains.harmony.config.ProverConfig",
    "trusting_period": "120h"
  }
}
