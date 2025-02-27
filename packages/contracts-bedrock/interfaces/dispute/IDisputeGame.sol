// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IInitializable } from "interfaces/dispute/IInitializable.sol";
import { Timestamp, GameStatus, GameType, Claim, Hash } from "src/dispute/lib/Types.sol";

interface IDisputeGame is IInitializable {
    /// @notice Emitted when the game is resolved.
    /// @param status The status of the game after resolution.
    /// @param rootClaim The root claim of the DisputeGame.
    /// @param l2BlockNumber The L2 block number at which the output root was generated.
    event Resolved(GameStatus indexed status, Claim indexed rootClaim, uint256 indexed l2BlockNumber);

    function createdAt() external view returns (Timestamp);
    function resolvedAt() external view returns (Timestamp);
    function status() external view returns (GameStatus);
    function gameType() external view returns (GameType gameType_);
    function gameCreator() external pure returns (address creator_);
    function rootClaim() external pure returns (Claim rootClaim_);
    function l1Head() external pure returns (Hash l1Head_);
    function l2BlockNumber() external pure returns (uint256 l2BlockNumber_);
    function extraData() external pure returns (bytes memory extraData_);
    function resolve() external returns (GameStatus status_);
    function gameData() external view returns (GameType gameType_, Claim rootClaim_, bytes memory extraData_);
    function wasRespectedGameTypeWhenCreated() external view returns (bool);
}
