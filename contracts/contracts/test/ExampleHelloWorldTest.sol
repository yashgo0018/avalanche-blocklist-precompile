//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "../ExampleHelloWorld.sol";
import "../interfaces/IHelloWorld.sol";
import "@avalabs/subnet-evm-contracts/contracts/test/AllowListTest.sol";

contract ExampleHelloWorldTest is AllowListTest {
    IHelloWorld public helloWorld = IHelloWorld(HELLO_WORLD_ADDRESS);

    function step_getDefaultHelloWorld() public {
        ExampleHelloWorld example = new ExampleHelloWorld();
        address exampleAddress = address(example);
        assertRole(helloWorld.readAllowList(exampleAddress), AllowList.Role.None);
        assertEq(example.sayHello(), "Hello World!");
    }

    function step_doesNotSetGreetingBeforeEnabled() public {
        ExampleHelloWorld example = new ExampleHelloWorld();
        address exampleAddress = address(example);
        assertRole(helloWorld.readAllowList(exampleAddress), AllowList.Role.None);
        try example.setGreeting("testing") {
            assertTrue(false, "setGreeting should fail");
        } catch {}
    }

    function step_setAndGetGreeting() public {
        ExampleHelloWorld example = new ExampleHelloWorld();
        address exampleAddress = address(example);
        assertRole(helloWorld.readAllowList(exampleAddress), AllowList.Role.None);
        helloWorld.setEnabled(exampleAddress);
        assertRole(
            helloWorld.readAllowList(exampleAddress),
            AllowList.Role.Enabled
        );
        string memory greeting = "testgreeting";
        example.setGreeting(greeting);
        assertEq(example.sayHello(), greeting);
    }
}
