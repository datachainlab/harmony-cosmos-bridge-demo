const IBCHost = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHost");
const IBCHandler = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHandler");
const TendermintLightClient = artifacts.require("@datachainlab/tendermint-sol/TendermintLightClient");
const ICS20TransferBank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20TransferBank");
const ICS20Bank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20Bank");

const PortTransfer = "transfer"
const TendermintLightClientType = "07-tendermint"

module.exports = async function (deployer) {
  const ibcHost = await IBCHost.deployed();
  const ibcHandler = await IBCHandler.deployed();
  const ics20Bank = await ICS20Bank.deployed();

  for(const f of [
    () => ibcHost.setIBCModule(IBCHandler.address),
    () => ibcHandler.bindPort(PortTransfer, ICS20TransferBank.address),
    () => ibcHandler.registerClient(TendermintLightClientType, TendermintLightClient.address),
    () => ics20Bank.setOperator(ICS20TransferBank.address),
  ]) {
    const result = await f().catch((err) => { throw err });
    console.log(result);
    if(!result.receipt.status) {
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
};
