import { Resolved as ResolvedEvent } from "../generated/templates/FaultDisputeGame/FaultDisputeGame"
import { Resolved } from "../generated/schema"
import { DisputeGameCreated } from "../generated/schema"

export function handleResolved(event: ResolvedEvent): void {
  let entity = new Resolved(
    event.transaction.hash.toHex() + "-" + event.logIndex.toString(),
  )
  entity.status = event.params.status
  entity.save()

  // update dispute game
  let disputeGame = DisputeGameCreated.load(event.address)
  if (disputeGame == null) {
    return
  }
  disputeGame.resolved = entity.id
  disputeGame.save()
}
