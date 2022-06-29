const ICS20Bank = artifacts.require("@hyperledger-labs/yui-ibc-solidity/ICS20Bank");
const SimpleToken = artifacts.require("@hyperledger-labs/yui-ibc-solidity/SimpleToken");

module.exports = async function (deployer, network, accounts) {
  const ics20Bank = await ICS20Bank.deployed();
  const token = await SimpleToken.deployed();

  for(const f of [
    () => token.approve(ICS20Bank.address, 1000000),
    () => ics20Bank.deposit(SimpleToken.address, 1000000, accounts[0])
  ]) {
    const result = await f().catch((err) => { throw err });
    console.log(result);
    if(!result.receipt.status) {
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
};
