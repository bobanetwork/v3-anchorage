import { DeployFunction, DeploymentSubmission } from 'hardhat-deploy/types'
import { ethers } from 'hardhat'
import { Contract, ContractFactory } from 'ethers'
import { DeterministicDeployer } from '../src/DeterministicDeployer'
import MockFeedRegistryJson from '../artifacts/contracts/test/mocks/MockFeedRegistry.sol/MockFeedRegistry.json'
import L2GovernanceERC20Json from '@eth-optimism/contracts-bedrock/artifacts/src/boba/L2GovernanceERC20.sol/L2GovernanceERC20.json'

let Factory__MockFeedRegistry: ContractFactory
let MockFeedRegistry: Contract
let Factory__L2GovernanceERC20: ContractFactory
let L2GovernanceERC20: Contract

const deployFn: DeployFunction = async (hre) => {
    // deploy these helpers only for the local devnet
    if ((hre as any).deployConfig.network === 'local') {
        Factory__MockFeedRegistry = new ContractFactory(
            MockFeedRegistryJson.abi,
            MockFeedRegistryJson.bytecode,
            (hre as any).deployConfig.deployer_l2
        )
        MockFeedRegistry = await Factory__MockFeedRegistry.deploy()
        await MockFeedRegistry.deployTransaction.wait()

        const MockFeedRegistryDeploymentSubmission: DeploymentSubmission = {
            address: MockFeedRegistry.address,
            abi: MockFeedRegistryJson.abi
          }
        await hre.deployments.save('MockFeedRegistry', MockFeedRegistryDeploymentSubmission)

        Factory__L2GovernanceERC20 = new ContractFactory(
            L2GovernanceERC20Json.abi,
            L2GovernanceERC20Json.bytecode,
            (hre as any).deployConfig.deployer_l2
        )
        L2GovernanceERC20 = await Factory__L2GovernanceERC20.deploy('0x4200000000000000000000000000000000000010', (hre as any).deployConfig.l1BobaTokenAddress, 'BOBA', 'BOBA', 18)
        await L2GovernanceERC20.deployTransaction.wait()

        const L2GovernanceERC20DeploymentSubmission: DeploymentSubmission = {
            address: L2GovernanceERC20.address,
            abi: L2GovernanceERC20Json.abi
          }
        await hre.deployments.save('L2GovernanceERC20', L2GovernanceERC20DeploymentSubmission)
    }
}

export default deployFn
deployFn.tags = ['Helper']