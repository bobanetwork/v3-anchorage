import { Wallet, providers, Contract } from 'ethers'

/* eslint-disable */
require('dotenv').config()

import hre, { ethers } from 'hardhat'

import AddressManagerJson from '../artifacts/contracts/utils/AddressManager.sol/AddressManager.json'

const main = async () => {
  console.log('Starting BOBA AA contracts deployment...')

  const network = process.env.NETWORK || 'local'

  const l1Provider = new providers.JsonRpcProvider(process.env.L1_NODE_WEB3_URL)
  const l2Provider = new providers.JsonRpcProvider(process.env.L2_NODE_WEB3_URL)

  const deployer_l1 = new Wallet(process.env.DEPLOYER_PRIVATE_KEY, l1Provider)
  const deployer_l2 = new Wallet(process.env.DEPLOYER_PRIVATE_KEY, l2Provider)

  const getAddressManager = (provider: any, addressManagerAddress: any) => {
    return new Contract(
      addressManagerAddress,
      AddressManagerJson.abi,
      provider
    )
  }
  let addressManagerAddress = process.env.ADDRESS_MANAGER_ADDRESS
  if (addressManagerAddress === ''|| typeof addressManagerAddress === 'undefined') {
    console.warn('ADDRESS_MANAGER_ADDRESS is not set, deploying AddressManager')
    const Factory__AddressManager = new ethers.ContractFactory(
      AddressManagerJson.abi,
      AddressManagerJson.bytecode,
      deployer_l1
    )
    const addressManager = await Factory__AddressManager.deploy()
    await addressManager.deployTransaction.wait()
    addressManagerAddress = addressManager.address
    console.log(`AddressManager deployed at ${addressManagerAddress}`)
  }
  console.log(
    `ADDRESS_MANAGER_ADDRESS was set to ${addressManagerAddress}`
  )

  const addressManager = getAddressManager(
    deployer_l1,
    addressManagerAddress
  )
  await hre.deployments.save('AddressManager', {
    address: addressManager.address,
    abi: AddressManagerJson.abi,
  })

  await hre.run('deploy', {
    l1Provider,
    l2Provider,
    deployer_l1,
    deployer_l2,
    addressManager,
    network,
    noCompile: process.env.NO_COMPILE ? true : false,
  })
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.log(
      JSON.stringify({ error: error.message, stack: error.stack }, null, 2)
    )
    process.exit(1)
  })
