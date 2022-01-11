const IBCHost = artifacts.require("IBCHost");
// TODO replace with Tendermint Client
const MockClient = artifacts.require("MockClient");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannel = artifacts.require("IBCChannel");
const IBCHandler = artifacts.require("IBCHandler");
const IBCMsgs = artifacts.require("IBCMsgs");
const IBCIdentifier = artifacts.require("IBCIdentifier");
const SimpleToken = artifacts.require("SimpleToken");
const ICS20TransferBank = artifacts.require("ICS20TransferBank");
const ICS20Bank = artifacts.require("ICS20Bank");

module.exports = function (deployer) {
  deployer.deploy(IBCIdentifier).then(function() {
    return deployer.link(IBCIdentifier, [IBCHost, IBCHandler]);
  });
  deployer.deploy(IBCMsgs).then(function() {
    return deployer.link(IBCMsgs, [IBCClient, IBCConnection, IBCChannel, IBCHandler]);
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
  deployer.deploy(MockClient);
  deployer.deploy(IBCHost).then(function() {
    return deployer.deploy(IBCHandler, IBCHost.address);
  });
  deployer.deploy(SimpleToken, "simple", "simple", 1000000);
  deployer.deploy(ICS20Bank).then(function() {
    return deployer.deploy(ICS20TransferBank, IBCHost.address, IBCHandler.address, ICS20Bank.address);
  });
};
