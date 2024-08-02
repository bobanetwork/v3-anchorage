import {
  DisputeGameCreated as DisputeGameCreatedEvent,
  ImplementationSet as ImplementationSetEvent,
  InitBondUpdated as InitBondUpdatedEvent,
  Initialized as InitializedEvent,
  OwnershipTransferred as OwnershipTransferredEvent,
} from "../generated/DisputeGameFactory/DisputeGameFactory"
import {
  DisputeGameCreated,
  ImplementationSet,
  InitBondUpdated,
  Initialized,
  OwnershipTransferred,
} from "../generated/schema"

export function handleDisputeGameCreated(event: DisputeGameCreatedEvent): void {
  let entity = new DisputeGameCreated(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.disputeProxy = event.params.disputeProxy
  entity.gameType = event.params.gameType
  entity.rootClaim = event.params.rootClaim

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleImplementationSet(event: ImplementationSetEvent): void {
  let entity = new ImplementationSet(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.impl = event.params.impl
  entity.gameType = event.params.gameType

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleInitBondUpdated(event: InitBondUpdatedEvent): void {
  let entity = new InitBondUpdated(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.gameType = event.params.gameType
  entity.newBond = event.params.newBond

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleInitialized(event: InitializedEvent): void {
  let entity = new Initialized(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.version = event.params.version

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleOwnershipTransferred(
  event: OwnershipTransferredEvent,
): void {
  let entity = new OwnershipTransferred(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.previousOwner = event.params.previousOwner
  entity.newOwner = event.params.newOwner

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
