import { DeployFunction, DeploymentSubmission } from 'hardhat-deploy/types'
import { ethers } from 'hardhat'
import { Contract, ContractFactory } from 'ethers'
import BobaVerifyingPaymasterJson from '../artifacts/contracts/samples/BobaVerifyingPaymaster.sol/BobaVerifyingPaymaster.json'
import { DeterministicDeployer } from '../src/DeterministicDeployer'

let Factory__BobaVerifyingPaymaster: ContractFactory
let BobaVerifyingPaymaster: Contract

const deployFn: DeployFunction = async (hre) => {
  Factory__BobaVerifyingPaymaster = new ContractFactory(
    BobaVerifyingPaymasterJson.abi,
    BobaVerifyingPaymasterJson.bytecode,
    (hre as any).deployConfig.deployer_l2
  )
  // use the sig verifying address
  const verifyingSignerAddress = (hre as any).deployConfig.deployer_l2.address
  console.log(`Verifying Signer is: ${verifyingSignerAddress}`)

  const entryPoint = await hre.deployments.getOrNull('EntryPoint')
  console.log(`EntryPoint is located at: ${entryPoint.address}`)
  const bobaDepositPaymaster = await hre.deployments.getOrNull('BobaDepositPaymaster')
  console.log(`BobaDepositPaymaster is located at: ${bobaDepositPaymaster.address}`)
  const bobaToken = (hre as any).deployConfig.l2BobaTokenAddress || (await hre.deployments.getOrNull('L2GovernanceERC20')).address
  console.log(bobaToken)
  console.log(`Boba is located at: ${bobaToken}`)
  BobaVerifyingPaymaster = await Factory__BobaVerifyingPaymaster.deploy(entryPoint.address, verifyingSignerAddress, bobaDepositPaymaster.address, bobaToken)
  await BobaVerifyingPaymaster.deployTransaction.wait()

  console.log('Boba Verifying Paymaster at', BobaVerifyingPaymaster.address)

  const BobaVerifyingPaymasterDeploymentSubmission: DeploymentSubmission = {
    address: BobaVerifyingPaymaster.address,
    abi: BobaVerifyingPaymasterJson.abi
  }
  await hre.deployments.save('BobaVerifyingPaymaster', BobaVerifyingPaymasterDeploymentSubmission)
}

export default deployFn
deployFn.tags = ['BobaVerifyingPaymaster']
