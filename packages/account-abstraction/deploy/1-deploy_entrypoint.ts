import { DeployFunction, DeploymentSubmission } from 'hardhat-deploy/types'
import { ethers } from 'hardhat'
import { Contract, ContractFactory } from 'ethers'
import EntryPointJson from '../artifacts/contracts/core/EntryPoint.sol/EntryPoint.json'
import { DeterministicDeployer } from '../src/DeterministicDeployer'

const sleep = async (ms: number): Promise<void> => {
  return new Promise<void>((resolve) => {
    setTimeout(() => {
      resolve(null)
    }, ms)
  })
}

const hexStringEquals = (stringA: string, stringB: string): boolean => {
  if (!ethers.utils.isHexString(stringA)) {
    throw new Error(`input is not a hex string: ${stringA}`)
  }

  if (!ethers.utils.isHexString(stringB)) {
    throw new Error(`input is not a hex string: ${stringB}`)
  }

  return stringA.toLowerCase() === stringB.toLowerCase()
}

const waitUntilTrue = async (
  check: () => Promise<boolean>,
  opts: {
    retries?: number
    delay?: number
  } = {}
) => {
  opts.retries = opts.retries || 100
  opts.delay = opts.delay || 5000

  let retries = 0
  while (!(await check())) {
    if (retries > opts.retries) {
      throw new Error(`check failed after ${opts.retries} attempts`)
    }
    retries++
    await sleep(opts.delay)
  }
}

let Factory__EntryPoint: ContractFactory

const deployFn: DeployFunction = async (hre) => {
  // for local det deployer funds
  if ((hre as any).deployConfig.network === 'local') {
    const txResponse = await (hre as any).deployConfig.deployer_l2.sendTransaction({to: '0xdDC61b272C2120D3e017886D647427d9e65603CF', value: '1000000000000000000'})
    await txResponse.wait()
  }

  Factory__EntryPoint = new ContractFactory(
    EntryPointJson.abi,
    EntryPointJson.bytecode,
    (hre as any).deployConfig.deployer_l2
  )
  const dep = new DeterministicDeployer((hre as any).deployConfig.l2Provider, (hre as any).deployConfig.deployer_l2, 'devnet')
  await dep.deployDeployer()
  const EntryPointAddress = await dep.deterministicDeploy(Factory__EntryPoint.bytecode)
  console.log('Deployed EntryPoint at', EntryPointAddress)

  const EntryPointDeploymentSubmission: DeploymentSubmission = {
    address: EntryPointAddress,
    abi: EntryPointJson.abi
  }
  await hre.deployments.save('EntryPoint', EntryPointDeploymentSubmission)
}

export default deployFn
deployFn.tags = ['EntryPoint']
