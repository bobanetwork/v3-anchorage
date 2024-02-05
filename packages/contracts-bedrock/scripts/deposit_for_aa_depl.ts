import { Wallet, providers, ethers } from 'ethers'

/* eslint-disable */
require('dotenv').config()

import hre from 'hardhat'

async function main() {
    const l1Provider = new providers.JsonRpcProvider(process.env.L1_RPC)
    const l2Provider = new providers.JsonRpcProvider(process.env.L2_RPC)
  
    const deployer_l1 = new Wallet(process.env.PRIVATE_KEY_DEPLOYER, l1Provider)
    const deployer_l2 = new Wallet(process.env.PRIVATE_KEY_DEPLOYER, l2Provider)

    const L1Balance = await deployer_l1.getBalance()
    const L2Balance = await deployer_l2.getBalance()
    console.log(`L1 Balance: ${ethers.utils.formatEther(L1Balance)}`)
    console.log(`L2 Balance: ${ethers.utils.formatEther(L2Balance)}`)
    console.log(`---------------------------------------------------`)
  
    const L1DepositAmount = ethers.utils.parseEther('10')
    if (L1Balance.lt(L1DepositAmount)) {
      console.error('Insufficient L1 balance')
      return
    }
  
    const L1Tx = await deployer_l1.sendTransaction({
      to: process.env.OPTIMISMPORTAL,
      value: L1DepositAmount,
    })
    await L1Tx.wait()

    while (true) {
        const postL2Balance = await deployer_l2.getBalance()
        if (!L2Balance.eq(postL2Balance)) {
          break
        }
        await new Promise((resolve) => setTimeout(resolve, 1000))
      }

    // deposit to create2deployer address, this is an address specifically used for local devnet
    const txResponse = await deployer_l2.sendTransaction({to: '0xdDC61b272C2120D3e017886D647427d9e65603CF', value: '1000000000000000000'})

    await txResponse.wait()

    console.log('Done!')
  }
  
  main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });