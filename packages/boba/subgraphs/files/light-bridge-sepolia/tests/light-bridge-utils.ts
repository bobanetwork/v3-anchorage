import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  AssetBalanceWithdrawn,
  AssetReceived,
  DisbursementFailed,
  DisbursementRetrySuccess,
  DisbursementSuccess,
  DisburserTransferred,
  MaxDepositAmountSet,
  MaxTransferAmountPerDaySet,
  MinDepositAmountSet,
  OwnershipTransferred,
  Paused,
  TokenSupported,
  Unpaused
} from "../generated/LightBridge/LightBridge"

export function createAssetBalanceWithdrawnEvent(
  token: Address,
  owner: Address,
  balance: BigInt
): AssetBalanceWithdrawn {
  let assetBalanceWithdrawnEvent = changetype<AssetBalanceWithdrawn>(
    newMockEvent()
  )

  assetBalanceWithdrawnEvent.parameters = new Array()

  assetBalanceWithdrawnEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  assetBalanceWithdrawnEvent.parameters.push(
    new ethereum.EventParam("owner", ethereum.Value.fromAddress(owner))
  )
  assetBalanceWithdrawnEvent.parameters.push(
    new ethereum.EventParam(
      "balance",
      ethereum.Value.fromUnsignedBigInt(balance)
    )
  )

  return assetBalanceWithdrawnEvent
}

export function createAssetReceivedEvent(
  token: Address,
  sourceChainId: BigInt,
  toChainId: BigInt,
  depositId: BigInt,
  emitter: Address,
  amount: BigInt
): AssetReceived {
  let assetReceivedEvent = changetype<AssetReceived>(newMockEvent())

  assetReceivedEvent.parameters = new Array()

  assetReceivedEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  assetReceivedEvent.parameters.push(
    new ethereum.EventParam(
      "sourceChainId",
      ethereum.Value.fromUnsignedBigInt(sourceChainId)
    )
  )
  assetReceivedEvent.parameters.push(
    new ethereum.EventParam(
      "toChainId",
      ethereum.Value.fromUnsignedBigInt(toChainId)
    )
  )
  assetReceivedEvent.parameters.push(
    new ethereum.EventParam(
      "depositId",
      ethereum.Value.fromUnsignedBigInt(depositId)
    )
  )
  assetReceivedEvent.parameters.push(
    new ethereum.EventParam("emitter", ethereum.Value.fromAddress(emitter))
  )
  assetReceivedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )

  return assetReceivedEvent
}

export function createDisbursementFailedEvent(
  depositId: BigInt,
  to: Address,
  amount: BigInt,
  sourceChainId: BigInt
): DisbursementFailed {
  let disbursementFailedEvent = changetype<DisbursementFailed>(newMockEvent())

  disbursementFailedEvent.parameters = new Array()

  disbursementFailedEvent.parameters.push(
    new ethereum.EventParam(
      "depositId",
      ethereum.Value.fromUnsignedBigInt(depositId)
    )
  )
  disbursementFailedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  disbursementFailedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  disbursementFailedEvent.parameters.push(
    new ethereum.EventParam(
      "sourceChainId",
      ethereum.Value.fromUnsignedBigInt(sourceChainId)
    )
  )

  return disbursementFailedEvent
}

export function createDisbursementRetrySuccessEvent(
  depositId: BigInt,
  to: Address,
  amount: BigInt,
  sourceChainId: BigInt
): DisbursementRetrySuccess {
  let disbursementRetrySuccessEvent = changetype<DisbursementRetrySuccess>(
    newMockEvent()
  )

  disbursementRetrySuccessEvent.parameters = new Array()

  disbursementRetrySuccessEvent.parameters.push(
    new ethereum.EventParam(
      "depositId",
      ethereum.Value.fromUnsignedBigInt(depositId)
    )
  )
  disbursementRetrySuccessEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  disbursementRetrySuccessEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  disbursementRetrySuccessEvent.parameters.push(
    new ethereum.EventParam(
      "sourceChainId",
      ethereum.Value.fromUnsignedBigInt(sourceChainId)
    )
  )

  return disbursementRetrySuccessEvent
}

export function createDisbursementSuccessEvent(
  depositId: BigInt,
  to: Address,
  token: Address,
  amount: BigInt,
  sourceChainId: BigInt
): DisbursementSuccess {
  let disbursementSuccessEvent = changetype<DisbursementSuccess>(newMockEvent())

  disbursementSuccessEvent.parameters = new Array()

  disbursementSuccessEvent.parameters.push(
    new ethereum.EventParam(
      "depositId",
      ethereum.Value.fromUnsignedBigInt(depositId)
    )
  )
  disbursementSuccessEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  disbursementSuccessEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  disbursementSuccessEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  disbursementSuccessEvent.parameters.push(
    new ethereum.EventParam(
      "sourceChainId",
      ethereum.Value.fromUnsignedBigInt(sourceChainId)
    )
  )

  return disbursementSuccessEvent
}

export function createDisburserTransferredEvent(
  newDisburser: Address
): DisburserTransferred {
  let disburserTransferredEvent = changetype<DisburserTransferred>(
    newMockEvent()
  )

  disburserTransferredEvent.parameters = new Array()

  disburserTransferredEvent.parameters.push(
    new ethereum.EventParam(
      "newDisburser",
      ethereum.Value.fromAddress(newDisburser)
    )
  )

  return disburserTransferredEvent
}

export function createMaxDepositAmountSetEvent(
  token: Address,
  toChainId: BigInt,
  previousAmount: BigInt,
  newAmount: BigInt
): MaxDepositAmountSet {
  let maxDepositAmountSetEvent = changetype<MaxDepositAmountSet>(newMockEvent())

  maxDepositAmountSetEvent.parameters = new Array()

  maxDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  maxDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam(
      "toChainId",
      ethereum.Value.fromUnsignedBigInt(toChainId)
    )
  )
  maxDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam(
      "previousAmount",
      ethereum.Value.fromUnsignedBigInt(previousAmount)
    )
  )
  maxDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam(
      "newAmount",
      ethereum.Value.fromUnsignedBigInt(newAmount)
    )
  )

  return maxDepositAmountSetEvent
}

export function createMaxTransferAmountPerDaySetEvent(
  token: Address,
  toChainId: BigInt,
  previousAmount: BigInt,
  newAmount: BigInt
): MaxTransferAmountPerDaySet {
  let maxTransferAmountPerDaySetEvent = changetype<MaxTransferAmountPerDaySet>(
    newMockEvent()
  )

  maxTransferAmountPerDaySetEvent.parameters = new Array()

  maxTransferAmountPerDaySetEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  maxTransferAmountPerDaySetEvent.parameters.push(
    new ethereum.EventParam(
      "toChainId",
      ethereum.Value.fromUnsignedBigInt(toChainId)
    )
  )
  maxTransferAmountPerDaySetEvent.parameters.push(
    new ethereum.EventParam(
      "previousAmount",
      ethereum.Value.fromUnsignedBigInt(previousAmount)
    )
  )
  maxTransferAmountPerDaySetEvent.parameters.push(
    new ethereum.EventParam(
      "newAmount",
      ethereum.Value.fromUnsignedBigInt(newAmount)
    )
  )

  return maxTransferAmountPerDaySetEvent
}

export function createMinDepositAmountSetEvent(
  token: Address,
  toChainId: BigInt,
  previousAmount: BigInt,
  newAmount: BigInt
): MinDepositAmountSet {
  let minDepositAmountSetEvent = changetype<MinDepositAmountSet>(newMockEvent())

  minDepositAmountSetEvent.parameters = new Array()

  minDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  minDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam(
      "toChainId",
      ethereum.Value.fromUnsignedBigInt(toChainId)
    )
  )
  minDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam(
      "previousAmount",
      ethereum.Value.fromUnsignedBigInt(previousAmount)
    )
  )
  minDepositAmountSetEvent.parameters.push(
    new ethereum.EventParam(
      "newAmount",
      ethereum.Value.fromUnsignedBigInt(newAmount)
    )
  )

  return minDepositAmountSetEvent
}

export function createOwnershipTransferredEvent(
  newOwner: Address
): OwnershipTransferred {
  let ownershipTransferredEvent = changetype<OwnershipTransferred>(
    newMockEvent()
  )

  ownershipTransferredEvent.parameters = new Array()

  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam("newOwner", ethereum.Value.fromAddress(newOwner))
  )

  return ownershipTransferredEvent
}

export function createPausedEvent(account: Address): Paused {
  let pausedEvent = changetype<Paused>(newMockEvent())

  pausedEvent.parameters = new Array()

  pausedEvent.parameters.push(
    new ethereum.EventParam("account", ethereum.Value.fromAddress(account))
  )

  return pausedEvent
}

export function createTokenSupportedEvent(
  token: Address,
  toChainId: BigInt,
  supported: boolean
): TokenSupported {
  let tokenSupportedEvent = changetype<TokenSupported>(newMockEvent())

  tokenSupportedEvent.parameters = new Array()

  tokenSupportedEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  tokenSupportedEvent.parameters.push(
    new ethereum.EventParam(
      "toChainId",
      ethereum.Value.fromUnsignedBigInt(toChainId)
    )
  )
  tokenSupportedEvent.parameters.push(
    new ethereum.EventParam("supported", ethereum.Value.fromBoolean(supported))
  )

  return tokenSupportedEvent
}

export function createUnpausedEvent(account: Address): Unpaused {
  let unpausedEvent = changetype<Unpaused>(newMockEvent())

  unpausedEvent.parameters = new Array()

  unpausedEvent.parameters.push(
    new ethereum.EventParam("account", ethereum.Value.fromAddress(account))
  )

  return unpausedEvent
}
