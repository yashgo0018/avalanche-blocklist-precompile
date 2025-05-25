//SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IValidatorSet {
    // Set [addr] to have the admin role over the precompile contract.
    function height() external view returns (uint256 height);

    // Set [addr] to be enabled on the precompile contract.
    function validator(address addr, string calldata reason) external;

    // Set [addr] to have the manager role over the precompile contract.
    function unblockAddress(address addr, string calldata reason) external;

    // Read the status of [addr].
    function readBlockList(address addr) external view returns (uint256 role);

    // Read the admin address
    function admin() external view returns (address addr);
}
