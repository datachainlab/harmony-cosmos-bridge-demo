const IBCHost = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHost");
const IBCClient = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCClient");
const IBCConnection = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCConnection");
const IBCChannel = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCChannel");
const IBCHandler = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHandler");
const IBCMsgs = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCMsgs");
const IBCIdentifier = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCIdentifier");
const TendermintLightClient = artifacts.require("@datachainlab/tendermint-sol/TendermintLightClient");
  // libs
const Bytes = artifacts.require("@datachainlab/tendermint-sol/Bytes");

const SimpleToken = artifacts.require("@hyperledger-labs/yui-ibc-solidity/SimpleToken");
const ICS20TransferBank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20TransferBank");
const ICS20Bank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20Bank");

module.exports = function (deployer) {
  deployer.deploy(IBCIdentifier).then(function() {
    return deployer.link(IBCIdentifier, [IBCHost, TendermintLightClient, IBCHandler]);
  });
  deployer.deploy(IBCMsgs).then(function() {
    return deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannel, IBCHandler, TendermintLightClient]);
  });
  deployer.deploy(IBCClient).then(function() {
    return deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannel]);
  });
  deployer.deploy(IBCConnection).then(function() {
    return deployer.link(IBCConnection, [IBCHandler, IBCChannel]);
  });
  deployer.deploy(IBCChannel).then(function() {
    return deployer.link(IBCChannel, [IBCHandler]);
  });
  deployer.deploy(Bytes);
  deployer.link(Bytes, TendermintLightClient);
  deployer.deploy(TendermintLightClient);

  deployer.deploy(IBCHost).then(function() {
    return deployer.deploy(IBCHandler, IBCHost.address);
  });

  deployer.deploy(SimpleToken, "simple", "simple", 1000000);
  deployer.deploy(ICS20Bank).then(function() {
    return deployer.deploy(ICS20TransferBank, IBCHost.address, IBCHandler.address, ICS20Bank.address);
  });
};
