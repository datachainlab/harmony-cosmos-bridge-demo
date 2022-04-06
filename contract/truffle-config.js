const { TruffleProvider } = require('@harmony-js/core')
//Local
const local_private_key = process.env.HARMONY_LOCAL_PRIVATE_KEY
const local_0_url = process.env.HARMONY_LOCAL_SHARD_0_URL
const local_1_url = process.env.HARMONY_LOCAL_SHARD_1_URL

//GAS - Currently using same GAS accross all environments
const gasLimit = process.env.HARMONY_GAS_LIMIT
const gasPrice = process.env.HARMONY_GAS_PRICE

module.exports = {
  networks: {
    local_shard_0: {
      network_id: '2',
      provider: () => {
        const truffleProvider = new TruffleProvider(
          local_0_url,
          {},
          { shardID: 0, chainId: 2 },
          { gasLimit: gasLimit, gasPrice: gasPrice},
        );
        const newAcc = truffleProvider.addByPrivateKey(local_private_key);
        truffleProvider.setSigner(newAcc);
        return truffleProvider;
      },
    },
    local_shard_1: {
      network_id: '2',
      provider: () => {
        const truffleProvider = new TruffleProvider(
          local_1_url,
          {},
          { shardID: 1, chainId: 2 },
          { gasLimit: gasLimit, gasPrice: gasPrice},
        );
        const newAcc = truffleProvider.addByPrivateKey(local_private_key);
        truffleProvider.setSigner(newAcc);
        return truffleProvider;
      },
    },
  },

  // Set default mocha options here, use special reporters etc.
  mocha: {
    // timeout: 100000
  },

  // Configure your compilers
  compilers: {
    solc: {
      version: "0.8.9",    // Fetch exact version from solc-bin (default: truffle's version)
      // docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
      settings: {          // See the solidity docs for advice about optimization and evmVersion
      optimizer: {
        enabled: true,
        runs: 1000
      },
      //  evmVersion: "byzantium"
      }
    }
  }
}
