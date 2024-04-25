import { DeployFunction, DeploymentSubmission } from 'hardhat-deploy/types'
import { Contract, ContractFactory, constants } from 'ethers'
import { registerBobaAddress } from './1-deploy_entrypoint'
import BobaDepositPaymasterJson from '../artifacts/contracts/samples/BobaDepositPaymaster.sol/BobaDepositPaymaster.json'
import MockFeedRegistry from '../artifacts/contracts/test/mocks/MockFeedRegistry.sol/MockFeedRegistry.json'

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
  let ethPriceOracle = await (hre as any).deployConfig.addressManager.getAddress('FeedRegistry')
  if (ethPriceOracle === constants.AddressZero) {
    console.warn('!!! WARNING: FeedRegistry not found in address manager, deploying a mock one')
    const Factory__MockFeedRegistry = new ContractFactory(
      MockFeedRegistry.abi,
      MockFeedRegistry.bytecode,
      (hre as any).deployConfig.deployer_l2
    )
    const mockFeedRegistry = await Factory__MockFeedRegistry.deploy()
    await mockFeedRegistry.deployTransaction.wait()
    ethPriceOracle = mockFeedRegistry.address
    console.log(`Mock FeedRegistry deployed at ${ethPriceOracle}`)
    const MockFeedRegistryDeploymentSubmission: DeploymentSubmission = {
      address: mockFeedRegistry.address,
      abi: MockFeedRegistry.abi
    }
    await hre.deployments.save('MockFeedRegistry', MockFeedRegistryDeploymentSubmission)
    await registerBobaAddress( (hre as any).deployConfig.addressManager, 'FeedRegistry', ethPriceOracle )
  }
  console.log(`Eth Price Oracle is located at: ${ethPriceOracle}`)
  const entryPointFromAM = await (hre as any).deployConfig.addressManager.getAddress('L2_Boba_EntryPoint')
  if (entryPoint.address.toLowerCase() === entryPointFromAM.toLowerCase()) {
    BobaDepositPaymaster = await Factory__BobaDepositPaymaster.deploy(entryPoint.address, ethPriceOracle)
    await BobaDepositPaymaster.deployTransaction.wait()

    console.log('Boba Deposit Paymaster at', BobaDepositPaymaster.address)

    const BobaDepositPaymasterDeploymentSubmission: DeploymentSubmission = {
      address: BobaDepositPaymaster.address,
      abi: BobaDepositPaymasterJson.abi
    }
    await hre.deployments.save('BobaDepositPaymaster', BobaDepositPaymasterDeploymentSubmission)

    await registerBobaAddress( (hre as any).deployConfig.addressManager, 'L2_BobaDepositPaymaster', BobaDepositPaymaster.address )
  }
}

export default deployFn
deployFn.tags = ['BobaDepositPaymaster']
