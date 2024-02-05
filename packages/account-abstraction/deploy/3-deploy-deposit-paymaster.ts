import { DeployFunction, DeploymentSubmission } from 'hardhat-deploy/types'
import { ethers } from 'hardhat'
import { Contract, ContractFactory } from 'ethers'
import BobaDepositPaymasterJson from '../artifacts/contracts/samples/BobaDepositPaymaster.sol/BobaDepositPaymaster.json'
import { DeterministicDeployer } from '../src/DeterministicDeployer'

let Factory__BobaDepositPaymaster: ContractFactory
let BobaDepositPaymaster: Contract

const deployFn: DeployFunction = async (hre) => {
  Factory__BobaDepositPaymaster = new ContractFactory(
    BobaDepositPaymasterJson.abi,
    BobaDepositPaymasterJson.bytecode,
    (hre as any).deployConfig.deployer_l2
  )
  const entryPoint = await hre.deployments.getOrNull('EntryPoint')
  console.log(`EntryPoint is located at: ${entryPoint.address}`)
  const ethPriceOracleAddress = (hre as any).deployConfig.ethPriceOracle || (await hre.deployments.getOrNull('MockFeedRegistry')).address
  console.log('ethPriceOracle', ethPriceOracleAddress)
  BobaDepositPaymaster = await Factory__BobaDepositPaymaster.deploy(entryPoint.address, ethPriceOracleAddress)
  await BobaDepositPaymaster.deployTransaction.wait()

  console.log('Boba Deposit Paymaster at', BobaDepositPaymaster.address)

  const BobaDepositPaymasterDeploymentSubmission: DeploymentSubmission = {
    address: BobaDepositPaymaster.address,
    abi: BobaDepositPaymasterJson.abi
  }
  await hre.deployments.save('BobaDepositPaymaster', BobaDepositPaymasterDeploymentSubmission)
}

export default deployFn
deployFn.tags = ['BobaDepositPaymaster']
