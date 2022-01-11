const IBCHost = artifacts.require("IBCHost");
const IBCHandler = artifacts.require("IBCHandler");
const MockClient = artifacts.require("MockClient");
const ICS20TransferBank = artifacts.require("ICS20TransferBank");
const ICS20Bank = artifacts.require("ICS20Bank");

const PortTransfer = "transfer"
const MockClientType = "mock-client"

module.exports = async function (deployer) {
  const ibcHost = await IBCHost.deployed();
  const ibcHandler = await IBCHandler.deployed();
  const ics20Bank = await ICS20Bank.deployed();

  for(const f of [
    () => ibcHost.setIBCModule(IBCHandler.address),
    () => ibcHandler.bindPort(PortTransfer, ICS20TransferBank.address),
    // TODO replace with Tendermint Client
    () => ibcHandler.registerClient(MockClientType, MockClient.address),
    () => ics20Bank.setOperator(ICS20TransferBank.address),
  ]) {
    const result = await f();
    if(!result.receipt.status) {
      console.log(result);
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
};
