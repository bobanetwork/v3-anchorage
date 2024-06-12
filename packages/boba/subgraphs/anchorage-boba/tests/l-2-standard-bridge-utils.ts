import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt, Bytes } from "@graphprotocol/graph-ts"
import {
  DepositFinalized,
  ERC20BridgeFinalized,
  ERC20BridgeInitiated,
  ETHBridgeFinalized,
  ETHBridgeInitiated,
  Initialized,
  WithdrawalInitiated
} from "../generated/L2StandardBridge/L2StandardBridge"

export function createDepositFinalizedEvent(
  l1Token: Address,
  l2Token: Address,
  from: Address,
  to: Address,
  amount: BigInt,
  extraData: Bytes
): DepositFinalized {
  let depositFinalizedEvent = changetype<DepositFinalized>(newMockEvent())

  depositFinalizedEvent.parameters = new Array()

  depositFinalizedEvent.parameters.push(
    new ethereum.EventParam("l1Token", ethereum.Value.fromAddress(l1Token))
  )
  depositFinalizedEvent.parameters.push(
    new ethereum.EventParam("l2Token", ethereum.Value.fromAddress(l2Token))
  )
  depositFinalizedEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  depositFinalizedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  depositFinalizedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  depositFinalizedEvent.parameters.push(
    new ethereum.EventParam("extraData", ethereum.Value.fromBytes(extraData))
  )

  return depositFinalizedEvent
}

export function createERC20BridgeFinalizedEvent(
  localToken: Address,
  remoteToken: Address,
  from: Address,
  to: Address,
  amount: BigInt,
  extraData: Bytes
): ERC20BridgeFinalized {
  let erc20BridgeFinalizedEvent = changetype<ERC20BridgeFinalized>(
    newMockEvent()
  )

  erc20BridgeFinalizedEvent.parameters = new Array()

  erc20BridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam(
      "localToken",
      ethereum.Value.fromAddress(localToken)
    )
  )
  erc20BridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam(
      "remoteToken",
      ethereum.Value.fromAddress(remoteToken)
    )
  )
  erc20BridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  erc20BridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  erc20BridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  erc20BridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("extraData", ethereum.Value.fromBytes(extraData))
  )

  return erc20BridgeFinalizedEvent
}

export function createERC20BridgeInitiatedEvent(
  localToken: Address,
  remoteToken: Address,
  from: Address,
  to: Address,
  amount: BigInt,
  extraData: Bytes
): ERC20BridgeInitiated {
  let erc20BridgeInitiatedEvent = changetype<ERC20BridgeInitiated>(
    newMockEvent()
  )

  erc20BridgeInitiatedEvent.parameters = new Array()

  erc20BridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam(
      "localToken",
      ethereum.Value.fromAddress(localToken)
    )
  )
  erc20BridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam(
      "remoteToken",
      ethereum.Value.fromAddress(remoteToken)
    )
  )
  erc20BridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  erc20BridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  erc20BridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  erc20BridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("extraData", ethereum.Value.fromBytes(extraData))
  )

  return erc20BridgeInitiatedEvent
}

export function createETHBridgeFinalizedEvent(
  from: Address,
  to: Address,
  amount: BigInt,
  extraData: Bytes
): ETHBridgeFinalized {
  let ethBridgeFinalizedEvent = changetype<ETHBridgeFinalized>(newMockEvent())

  ethBridgeFinalizedEvent.parameters = new Array()

  ethBridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  ethBridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  ethBridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  ethBridgeFinalizedEvent.parameters.push(
    new ethereum.EventParam("extraData", ethereum.Value.fromBytes(extraData))
  )

  return ethBridgeFinalizedEvent
}

export function createETHBridgeInitiatedEvent(
  from: Address,
  to: Address,
  amount: BigInt,
  extraData: Bytes
): ETHBridgeInitiated {
  let ethBridgeInitiatedEvent = changetype<ETHBridgeInitiated>(newMockEvent())

  ethBridgeInitiatedEvent.parameters = new Array()

  ethBridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  ethBridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  ethBridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  ethBridgeInitiatedEvent.parameters.push(
    new ethereum.EventParam("extraData", ethereum.Value.fromBytes(extraData))
  )

  return ethBridgeInitiatedEvent
}

export function createInitializedEvent(version: i32): Initialized {
  let initializedEvent = changetype<Initialized>(newMockEvent())

  initializedEvent.parameters = new Array()

  initializedEvent.parameters.push(
    new ethereum.EventParam(
      "version",
      ethereum.Value.fromUnsignedBigInt(BigInt.fromI32(version))
    )
  )

  return initializedEvent
}

export function createWithdrawalInitiatedEvent(
  l1Token: Address,
  l2Token: Address,
  from: Address,
  to: Address,
  amount: BigInt,
  extraData: Bytes
): WithdrawalInitiated {
  let withdrawalInitiatedEvent = changetype<WithdrawalInitiated>(newMockEvent())

  withdrawalInitiatedEvent.parameters = new Array()

  withdrawalInitiatedEvent.parameters.push(
    new ethereum.EventParam("l1Token", ethereum.Value.fromAddress(l1Token))
  )
  withdrawalInitiatedEvent.parameters.push(
    new ethereum.EventParam("l2Token", ethereum.Value.fromAddress(l2Token))
  )
  withdrawalInitiatedEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  withdrawalInitiatedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  withdrawalInitiatedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  withdrawalInitiatedEvent.parameters.push(
    new ethereum.EventParam("extraData", ethereum.Value.fromBytes(extraData))
  )

  return withdrawalInitiatedEvent
}
