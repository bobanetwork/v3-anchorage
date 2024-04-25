import { Wallet, providers, Contract } from 'ethers'

/* eslint-disable */
require('dotenv').config()

import hre, { ethers } from 'hardhat'

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
      ['function getAddress(string) external view returns (address)'],
      provider
    )
  }

  console.log(
    `ADDRESS_MANAGER_ADDRESS was set to ${process.env.ADDRESS_MANAGER_ADDRESS}`
  )
  const addressManager = getAddressManager(
    deployer_l1,
    process.env.ADDRESS_MANAGER_ADDRESS
  )

  let proxyAdminAddress = await addressManager.getAddress('ProxyAdmin')
  if (proxyAdminAddress === ethers.constants.AddressZero) {
    proxyAdminAddress = process.env.PROXY_ADMIN_ADDRESS
    if (proxyAdminAddress === ethers.constants.AddressZero) {
      throw new Error('ProxyAdmin address not found in address manager or in env')
    }
  }
  const proxyAdmin = new Contract(
    proxyAdminAddress,
    new ethers.utils.Interface(['function setAddress(string,address) external','function owner() external view returns (address)']),
    deployer_l1
  )
  console.log(await proxyAdmin.owner())

  let l1MessengerAddress = await addressManager.getAddress(
    'Proxy__L1CrossDomainMessenger'
  )
  if (l1MessengerAddress === ethers.constants.AddressZero) {
    l1MessengerAddress = process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS
    if (l1MessengerAddress === ethers.constants.AddressZero) {
      throw new Error(
        'L1CrossDomainMessenger address not found in address manager or in env'
      )
    }
  }

  let l2MessengerAddress = await addressManager.getAddress(
    'L2CrossDomainMessenger'
  )
  if (l2MessengerAddress === ethers.constants.AddressZero) {
    l2MessengerAddress = process.env.L2_CROSS_DOMAIN_MESSENGER_ADDRESS
    if (l2MessengerAddress === ethers.constants.AddressZero) {
      console.warn(`L2CrossDomainMessenger is set to 0x4200000000000000000000000000000000000007`)
      l2MessengerAddress = '0x4200000000000000000000000000000000000007'
    }
  }

  let L1StandardBridgeAddress = await addressManager.getAddress(
    'Proxy__L1StandardBridge'
  )
  if (L1StandardBridgeAddress === ethers.constants.AddressZero) {
    L1StandardBridgeAddress = process.env.L1_STANDARD_BRIDGE_ADDRESS
    if (L1StandardBridgeAddress === ethers.constants.AddressZero) {
      throw new Error(
        'L1StandardBridge address not found in address manager or in env'
      )
    }
  }
  const L1StandardBridge = new Contract(
    L1StandardBridgeAddress,
    new ethers.utils.Interface(['function otherBridge() external view returns (address)']),
    deployer_l1
  )

  const L2StandardBridgeAddress = await L1StandardBridge.otherBridge()

  await hre.run('deploy', {
    proxyAdminAddress,
    l1MessengerAddress,
    l2MessengerAddress,
    L1StandardBridgeAddress,
    L2StandardBridgeAddress,
    l1Provider,
    l2Provider,
    deployer_l1,
    deployer_l2,
    proxyAdmin,
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
