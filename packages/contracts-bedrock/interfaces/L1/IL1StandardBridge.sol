// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IStandardBridge } from "interfaces/universal/IStandardBridge.sol";
import { ICrossDomainMessenger } from "interfaces/universal/ICrossDomainMessenger.sol";
import { ISuperchainConfig } from "interfaces/L1/ISuperchainConfig.sol";

interface IL1StandardBridge is IStandardBridge {
    event ERC20DepositInitiated(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes extraData
    );
    event ERC20WithdrawalFinalized(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes extraData
    );
    event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData);
    event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData);

    function depositERC20(
        address _l1Token,
        address _l2Token,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes memory _extraData
    )
        external;
    function depositERC20To(
        address _l1Token,
        address _l2Token,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes memory _extraData
    )
        external;
    function depositETH(uint32 _minGasLimit, bytes memory _extraData) external payable;
    function depositETHTo(address _to, uint32 _minGasLimit, bytes memory _extraData) external payable;
    function finalizeERC20Withdrawal(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _extraData
    )
        external;
    function finalizeETHWithdrawal(
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _extraData
    )
        external
        payable;
    function initialize(
        ICrossDomainMessenger _messenger,
        ISuperchainConfig _superchainConfig
    )
        external;
    function l2TokenBridge() external view returns (address);
    function superchainConfig() external view returns (ISuperchainConfig);
    function version() external view returns (string memory);

    function __constructor__() external;
}
