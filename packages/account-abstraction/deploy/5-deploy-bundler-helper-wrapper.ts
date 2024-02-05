import { DeployFunction, DeploymentSubmission } from 'hardhat-deploy/types'
import { ethers } from 'hardhat'
import { Contract, ContractFactory } from 'ethers'
import EntryPointWrapperJson from '../artifacts/contracts/bundler/EntryPointWrapper.sol/EntryPointWrapper.json'
import { DeterministicDeployer } from '../src/DeterministicDeployer'

let Factory__EntryPointWrapper: ContractFactory

const deployFn: DeployFunction = async (hre) => {
  Factory__EntryPointWrapper = new ContractFactory(
    EntryPointWrapperJson.abi,
    EntryPointWrapperJson.bytecode,
    (hre as any).deployConfig.deployer_l2
  )
  const entryPoint = await hre.deployments.getOrNull('EntryPoint')
  console.log(`EntryPoint is located at: ${entryPoint.address}`)
    const entryPointWrapperConstructorArgs = ethers.utils.defaultAbiCoder.encode(
      ["address"],
      [entryPoint.address]
    )
    const entryPointWrapperCreationCode = ethers.utils.solidityPack(
      ["bytes", "bytes"],
      [Factory__EntryPointWrapper.bytecode, entryPointWrapperConstructorArgs]
    )
    const dep = new DeterministicDeployer((hre as any).deployConfig.l2Provider, (hre as any).deployConfig.deployer_l2, 'devnet')
    const EntryPointWrapperAddress = await dep.deterministicDeploy(entryPointWrapperCreationCode)
    console.log('EntryPoint Wrapper at', EntryPointWrapperAddress)

    const entryPointWrapperDeploymentSubmission: DeploymentSubmission = {
      address: EntryPointWrapperAddress,
      abi: EntryPointWrapperJson.abi
    }
    await hre.deployments.save('EntryPointWrapper', entryPointWrapperDeploymentSubmission)
}

export default deployFn
deployFn.tags = ['EntryPointWrapper']
