import { Wallet, providers } from 'ethers'

/* eslint-disable */
require('dotenv').config()

import hre from 'hardhat'

const main = async () => {
  console.log('Starting BOBA AA contracts deployment...')

  const network = process.env.NETWORK || 'local'

  const l1Provider = new providers.JsonRpcProvider(process.env.L1_NODE_WEB3_URL)
  const l2Provider = new providers.JsonRpcProvider(process.env.L2_NODE_WEB3_URL)

  const deployer_l1 = new Wallet(process.env.DEPLOYER_PRIVATE_KEY, l1Provider)
  const deployer_l2 = new Wallet(process.env.DEPLOYER_PRIVATE_KEY, l2Provider)

  console.log(
    `L1_BOBA_TOKEN_ADDRESS was set to ${process.env.L1_BOBA_TOKEN_ADDRESS}`
  )

  await hre.run('deploy', {
    l1BobaTokenAddress: process.env.L1_BOBA_TOKEN_ADDRESS,
    l2BobaTokenAddress: process.env.L2_BOBA_TOKEN_ADDRESS,
    ethPriceOracle: process.env.ETH_PRICE_ORACLE,
    l1Provider,
    l2Provider,
    deployer_l1,
    deployer_l2,
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
