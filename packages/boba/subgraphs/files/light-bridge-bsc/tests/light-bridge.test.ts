import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt } from "@graphprotocol/graph-ts"
import { AssetBalanceWithdrawn } from "../generated/schema"
import { AssetBalanceWithdrawn as AssetBalanceWithdrawnEvent } from "../generated/LightBridge/LightBridge"
import { handleAssetBalanceWithdrawn } from "../src/light-bridge"
import { createAssetBalanceWithdrawnEvent } from "./light-bridge-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let token = Address.fromString("0x0000000000000000000000000000000000000001")
    let owner = Address.fromString("0x0000000000000000000000000000000000000001")
    let balance = BigInt.fromI32(234)
    let newAssetBalanceWithdrawnEvent = createAssetBalanceWithdrawnEvent(
      token,
      owner,
      balance
    )
    handleAssetBalanceWithdrawn(newAssetBalanceWithdrawnEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("AssetBalanceWithdrawn created and stored", () => {
    assert.entityCount("AssetBalanceWithdrawn", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "AssetBalanceWithdrawn",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "token",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "AssetBalanceWithdrawn",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "owner",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "AssetBalanceWithdrawn",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "balance",
      "234"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
