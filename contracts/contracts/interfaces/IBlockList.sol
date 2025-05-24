//SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IBlockList {
    event AddressBlocked(address indexed account, string reason);
    event AddressUnblocked(address indexed account, string reason);
    event AdminChanged(address indexed account);

    // Set [addr] to have the admin role over the precompile contract.
    function changeAdmin(address addr) external;

    // Set [addr] to be enabled on the precompile contract.
    function blockAddress(address addr, string calldata reason) external;

    // Set [addr] to have the manager role over the precompile contract.
    function unblockAddress(address addr, string calldata reason) external;

    // Read the status of [addr].
    function readBlockList(address addr) external view returns (uint256 role);

    // Read the admin address
    function admin() external view returns (address addr);
}
