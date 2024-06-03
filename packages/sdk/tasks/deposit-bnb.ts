import { promises as fs } from 'fs'

import { task, types } from 'hardhat/config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'
import { Contract, providers, utils, constants } from 'ethers'
import { predeploys, sleep, hashWithdrawal } from '@eth-optimism/core-utils'
import Artifact__L2ToL1MessagePasser from '../src/forge-artifacts/L2ToL1MessagePasser.json'
import Artifact__L2StandardBridge from '../src/forge-artifacts/L2StandardBridge.json'
import Artifact__OptimismPortal from '../src/forge-artifacts/OptimismPortal.json'
import Artifact__L1StandardBridge from '../src/forge-artifacts/L1StandardBridge.json'
import Artifact__L2OutputOracle from '../src/forge-artifacts/L2OutputOracle.json'
import Artifact_BOBA from '../src/forge-artifacts/BOBA.json'

import {
  CONTRACT_ADDRESSES,
  OEContractsLike,
  DEFAULT_L2_CONTRACT_ADDRESSES,
} from '../src'

// Bridge a pre-deployed ERC20 from L1 to L2
task('deposit-bnb', 'Deposits BOBA and BNB onto L2.')
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL.',
    'http://localhost:9545',
    types.string
  )
  .addParam(
    'opNodeProviderUrl',
    'op-node provider URL',
    'http://localhost:7545',
    types.string
  )
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .setAction(async (args, hre) => {
    const signers = await hre.ethers.getSigners()
    if (signers.length === 0) {
      throw new Error('No configured signers')
    }
    // Use the first configured signer for simplicity
    const signer = signers[0]
    const address = await signer.getAddress()
    console.log(`Using signer ${address}`)

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await signer.getBalance()
    if (balance.eq(0)) {
      throw new Error('Signer has no balance')
    }

    const l1Provider = new providers.StaticJsonRpcProvider(args.l1ProviderUrl)

    const l1Signer = new hre.ethers.Wallet(
      hre.network.config.accounts[0],
      l1Provider
    )

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)

    const l2Signer = new hre.ethers.Wallet(
      hre.network.config.accounts[0],
      l2Provider
    )

    const l2ChainId = await l2Signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[l2ChainId]

    let l1BobaTokenAddress = ''
    const l2BobaTokenAddress = '0x4200000000000000000000000000000000000006'
    const l2BnbTokenAddress = '0x4200000000000000000000000000000000000023'

    if (args.l1ContractsJsonPath) {
      const data = await fs.readFile(args.l1ContractsJsonPath)
      const json = JSON.parse(data.toString())
      contractAddrs = {
        l1: {
          AddressManager: json.AddressManager,
          L1CrossDomainMessenger: json.L1CrossDomainMessengerProxy,
          L1StandardBridge: json.L1StandardBridgeProxy,
          StateCommitmentChain: constants.AddressZero,
          CanonicalTransactionChain: constants.AddressZero,
          BondManager: constants.AddressZero,
          OptimismPortal: json.OptimismPortalProxy,
          L2OutputOracle: json.L2OutputOracleProxy,
          OptimismPortal2: json.OptimismPortalProxy,
          DisputeGameFactory: json?.DisputeGameFactoryProxy,
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      } as OEContractsLike

      l1BobaTokenAddress = json.BOBA
    }

    console.log(`OptimismPortal: ${contractAddrs.l1.OptimismPortal}`)
    const OptimismPortal = new hre.ethers.Contract(
      contractAddrs.l1.OptimismPortal,
      Artifact__OptimismPortal.abi,
      signer
    )

    console.log(
      `L1CrossDomainMessenger: ${contractAddrs.l1.L1CrossDomainMessenger}`
    )

    console.log(`L1StandardBridge: ${contractAddrs.l1.L1StandardBridge}`)
    const L2OutputOracle = new hre.ethers.Contract(
      contractAddrs.l1.L2OutputOracle,
      Artifact__L2OutputOracle.abi,
      signer
    )

    const L2ToL1MessagePasser = new hre.ethers.Contract(
      predeploys.L2ToL1MessagePasser,
      Artifact__L2ToL1MessagePasser.abi,
      l2Signer
    )

    const L2StandardBridge = new hre.ethers.Contract(
      predeploys.L2StandardBridge,
      Artifact__L2StandardBridge.abi,
      l2Signer
    )

    const params = await OptimismPortal.params()
    console.log('Intial OptimismPortal.params:')
    console.log(params)

    const L1BobaToken = new Contract(
      l1BobaTokenAddress,
      Artifact_BOBA.abi,
      l1Signer
    )

    const L2BnbToken = new Contract(
      l2BnbTokenAddress,
      Artifact_BOBA.abi,
      l2Signer
    )

    let preL2BobaBalance = await l2Provider.getBalance(l2Signer.address)
    console.log(`Approving BOBA for deposit`)
    const approvalBobaTx = await L1BobaToken.approve(
      OptimismPortal.address,
      hre.ethers.constants.MaxUint256
    )
    await approvalBobaTx.wait()
    console.log('BOBA approved')

    console.log('Depositing BOBA to L2')
    const depositBobaTx = await OptimismPortal.depositERC20Transaction(
      l2Signer.address,
      utils.parseEther('10'),
      0,
      99999,
      false,
      '0x'
    )
    await depositBobaTx.wait()
    console.log(`ERC20 deposited - ${depositBobaTx.hash}`)

    console.log('Checking to make sure deposit was successful')
    // Deposit might get reorged, wait and also log for reorgs.
    let prevBlockHash: string = ''
    for (let i = 0; i < 12; i++) {
      const messageReceipt = await signer.provider!.getTransactionReceipt(
        depositBobaTx.hash
      )
      if (messageReceipt.status !== 1) {
        console.log(`Deposit failed, retrying...`)
      }

      // Wait for stability, we want some amount of time after any reorg
      if (prevBlockHash !== '' && messageReceipt.blockHash !== prevBlockHash) {
        console.log(
          `Block hash changed from ${prevBlockHash} to ${messageReceipt.blockHash}`
        )
        i = 0
      } else if (prevBlockHash !== '') {
        console.log(`No reorg detected: ${i}`)
      }

      prevBlockHash = messageReceipt.blockHash
      await sleep(1000)
    }
    console.log(`Deposit BOBA confirmed`)

    let postL2BobaBalance = await l2Provider.getBalance(l2Signer.address)
    if (!postL2BobaBalance.eq(preL2BobaBalance.add(utils.parseEther('10')))) {
      throw new Error(
        `bad deposit. recipient balance on L2: ${utils.formatEther(
          postL2BobaBalance
        )}`
      )
    }
    console.log(`Deposit BOBA success`)

    let preL2BnbBalance = await L2BnbToken.balanceOf(l2Signer.address)
    const depositBnbTx = await l1Signer.sendTransaction({
      to: OptimismPortal.address,
      value: utils.parseEther('10'),
    })
    await depositBnbTx.wait()
    console.log(`BNB deposited - ${depositBnbTx.hash}`)

    console.log('Checking to make sure deposit was successful')
    // Deposit might get reorged, wait and also log for reorgs.
    prevBlockHash = ''
    for (let i = 0; i < 12; i++) {
      const messageReceipt = await signer.provider!.getTransactionReceipt(
        depositBnbTx.hash
      )
      if (messageReceipt.status !== 1) {
        console.log(`Deposit failed, retrying...`)
      }

      // Wait for stability, we want some amount of time after any reorg
      if (prevBlockHash !== '' && messageReceipt.blockHash !== prevBlockHash) {
        console.log(
          `Block hash changed from ${prevBlockHash} to ${messageReceipt.blockHash}`
        )
        i = 0
      } else if (prevBlockHash !== '') {
        console.log(`No reorg detected: ${i}`)
      }

      prevBlockHash = messageReceipt.blockHash
      await sleep(1000)
    }
    console.log(`Deposit confirmed`)

    let postL2BnbBalance = await L2BnbToken.balanceOf(l2Signer.address)
    if (!postL2BnbBalance.eq(preL2BnbBalance.add(utils.parseEther('10')))) {
      throw new Error(
        `bad deposit. recipient balance on L2: ${utils.formatEther(
          postL2BnbBalance
        )}`
      )
    }
    console.log(`Deposit BNB success`)

    const finalizeWithdrawal = async (blockNumber: number) => {
      let latestSubmittedBlock = await L2OutputOracle.latestBlockNumber()

      while (latestSubmittedBlock < blockNumber) {
        console.log(
          `Waiting for block ${blockNumber} to be submitted, currently at ${latestSubmittedBlock}`
        )
        await sleep(5000)
        latestSubmittedBlock = await L2OutputOracle.latestBlockNumber()
      }

      const l2OutputIndex = await L2OutputOracle.getL2OutputIndexAfter(
        blockNumber
      )
      const proposal = await L2OutputOracle.getL2Output(l2OutputIndex)
      const l2BlockNumber = proposal.l2BlockNumber
      const outputRoot = proposal.outputRoot
      const block = await l2Provider.send('eth_getBlockByNumber', [
        l2BlockNumber.toNumber(),
        false,
      ])

      const messagePassedLogs = await L2ToL1MessagePasser.queryFilter(
        L2ToL1MessagePasser.filters.MessagePassed(),
        blockNumber,
        blockNumber
      )
      if (messagePassedLogs.length === 0) {
        throw new Error('No message passed logs found')
      }
      if (messagePassedLogs.length > 1) {
        throw new Error('More than one message passed log found')
      }
      const hash = hashWithdrawal(
        messagePassedLogs[0].args?.nonce,
        messagePassedLogs[0].args?.sender,
        messagePassedLogs[0].args?.target,
        messagePassedLogs[0].args?.value,
        messagePassedLogs[0].args?.gasLimit,
        messagePassedLogs[0].args?.data
      )
      const messageSlot = utils.keccak256(
        utils.defaultAbiCoder.encode(
          ['bytes32', 'uint256'],
          [hash, constants.HashZero]
        )
      )
      const withdrawalProof = await l2Provider.send('eth_getProof', [
        L2ToL1MessagePasser.address,
        [messageSlot],
        l2BlockNumber.toNumber(),
      ])

      const computedHash = utils.keccak256(
        utils.defaultAbiCoder.encode(
          ['bytes32', 'bytes32', 'bytes32', 'bytes32'],
          [
            constants.HashZero,
            block.stateRoot,
            withdrawalProof.storageHash,
            block.hash,
          ]
        )
      )

      if (computedHash !== outputRoot) {
        throw new Error('Output root mismatch')
      }

      const submitProofTx = await OptimismPortal.proveWithdrawalTransaction(
        [
          messagePassedLogs[0].args?.nonce,
          messagePassedLogs[0].args?.sender,
          messagePassedLogs[0].args?.target,
          messagePassedLogs[0].args?.value,
          messagePassedLogs[0].args?.gasLimit,
          messagePassedLogs[0].args?.data,
        ],
        l2OutputIndex,
        [
          constants.HashZero,
          block.stateRoot,
          withdrawalProof.storageHash,
          block.hash,
        ],
        withdrawalProof.storageProof[0].proof
      )
      await submitProofTx.wait()

      const gasEstimationFinalSubmit = async () => {
        const gasEstimation =
          await OptimismPortal.estimateGas.finalizeWithdrawalTransaction([
            messagePassedLogs[0].args?.nonce,
            messagePassedLogs[0].args?.sender,
            messagePassedLogs[0].args?.target,
            messagePassedLogs[0].args?.value,
            messagePassedLogs[0].args?.gasLimit,
            messagePassedLogs[0].args?.data,
          ])
        console.log(
          'Gas estimation for finalizeWithdrawalTransaction',
          gasEstimation.toString()
        )
      }

      let err = true
      while (err) {
        try {
          await gasEstimationFinalSubmit()
          err = false
        } catch (e) {
          console.log(
            'Failed to get gas estimation for finalizeWithdrawalTransaction'
          )
          await sleep(12000)
        }
      }
      const claimTokenTx = await OptimismPortal.finalizeWithdrawalTransaction([
        messagePassedLogs[0].args?.nonce,
        messagePassedLogs[0].args?.sender,
        messagePassedLogs[0].args?.target,
        messagePassedLogs[0].args?.value,
        messagePassedLogs[0].args?.gasLimit,
        messagePassedLogs[0].args?.data,
      ])
      await claimTokenTx.wait()
      console.log(`Withdrawal finalized - ${claimTokenTx.hash}`)
    }

    console.log('Starting BOBA withdrawal')
    const preL2BobaBalance = await l2Provider.getBalance(l2Signer.address)
    const preBobaBalance = await L1BobaToken.balanceOf(l1Signer.address)
    const withdrawalBobaTx = await l2Signer.sendTransaction({
      to: L2ToL1MessagePasser.address,
      value: utils.parseEther('1'),
    })
    const withdrawalBobaReceipt = await withdrawalBobaTx.wait()

    console.log('Starting BNB withdrawal')
    const opBalanceAfter = await signer!.provider!.getBalance(
      OptimismPortal.address
    )
    const approveBnbTx = await L2BnbToken.approve(
      L2StandardBridge.address,
      utils.parseEther('1')
    )
    await approveBnbTx.wait()
    const withdrawBnbTx = await L2StandardBridge.withdraw(
      l2BnbTokenAddress,
      utils.parseEther('1'),
      99999,
      '0x'
    )
    const withdrawalBnbReceipt = await withdrawBnbTx.wait()

    await finalizeWithdrawal(withdrawalBobaReceipt.blockNumber)

    const postBobaBalance = await L1BobaToken.balanceOf(l1Signer.address)
    const expectedBobaBalance = preBobaBalance.add(utils.parseEther('1'))
    if (!expectedBobaBalance.eq(postBobaBalance)) {
      throw new Error(
        `BOBA Balance mismatch, expected: ${expectedBobaBalance}, actual: ${postBobaBalance}`
      )
    }
    console.log('BOBA Withdrawal success')

    await finalizeWithdrawal(withdrawalBnbReceipt.blockNumber)

    const opBalanceFinally = await signer!.provider!.getBalance(
      OptimismPortal.address
    )

    const expectedopBalanceFinally = opBalanceFinally.add(utils.parseEther('1'))
    if (!expectedopBalanceFinally.eq(opBalanceAfter)) {
      throw new Error(
        `BNB Balance mismatch, expected: ${expectedopBalanceFinally}, actual: ${opBalanceAfter}`
      )
    }
    console.log('BNB Withdrawal success')
  })
