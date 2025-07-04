// Code generated
// This file is a generated precompile contract test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package blocklist

import (
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/subnet-evm/core/extstate"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/stretchr/testify/require"
)

var (
	_ = vm.ErrOutOfGas
	_ = big.NewInt
	_ = common.Big0
	_ = require.New
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
var (
	tests = map[string]testutils.PrecompileTest{
		"insufficient gas for admin should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				input, err := PackAdmin()
				require.NoError(t, err)
				return input
			},
			SuppliedGas: AdminGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vm.ErrOutOfGas.Error(),
		},
		"readOnly blockAddress should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := BlockAddressInput{}
				input, err := PackBlockAddress(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: BlockAddressGasCost,
			ReadOnly:    true,
			ExpectedErr: vm.ErrWriteProtection.Error(),
		},
		"insufficient gas for blockAddress should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := BlockAddressInput{}
				input, err := PackBlockAddress(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: BlockAddressGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vm.ErrOutOfGas.Error(),
		},
		"readOnly changeAdmin should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput common.Address
				testInput = common.Address{}
				input, err := PackChangeAdmin(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: ChangeAdminGasCost,
			ReadOnly:    true,
			ExpectedErr: vm.ErrWriteProtection.Error(),
		},
		"insufficient gas for changeAdmin should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput common.Address
				testInput = common.Address{}
				input, err := PackChangeAdmin(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: ChangeAdminGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vm.ErrOutOfGas.Error(),
		},
		"insufficient gas for readBlockList should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput common.Address
				testInput = common.Address{}
				input, err := PackReadBlockList(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: ReadBlockListGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vm.ErrOutOfGas.Error(),
		},
		"readOnly unblockAddress should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := UnblockAddressInput{}
				input, err := PackUnblockAddress(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: UnblockAddressGasCost,
			ReadOnly:    true,
			ExpectedErr: vm.ErrWriteProtection.Error(),
		},
		"insufficient gas for unblockAddress should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := UnblockAddressInput{}
				input, err := PackUnblockAddress(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: UnblockAddressGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vm.ErrOutOfGas.Error(),
		},
	}
)

// TestBlockListRun tests the Run function of the precompile contract.
func TestBlockListRun(t *testing.T) {
	// Run tests.
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.Run(t, Module, extstate.NewTestStateDB(t))
		})
	}
}

// TestPackUnpackAddressBlockedEventData tests the Pack/UnpackAddressBlockedEventData.
func TestPackUnpackAddressBlockedEventData(t *testing.T) {
	// CUSTOM CODE STARTS HERE
	// set test inputs with proper values here
	var accountInput common.Address = common.Address{}

	dataInput := AddressBlockedEventData{
		Reason: "",
	}

	_, data, err := PackAddressBlockedEvent(
		accountInput,
		dataInput,
	)
	require.NoError(t, err)

	unpacked, err := UnpackAddressBlockedEventData(data)
	require.NoError(t, err)
	require.Equal(t, dataInput, unpacked)
}

// TestPackUnpackAddressUnblockedEventData tests the Pack/UnpackAddressUnblockedEventData.
func TestPackUnpackAddressUnblockedEventData(t *testing.T) {
	// CUSTOM CODE STARTS HERE
	// set test inputs with proper values here
	var accountInput common.Address = common.Address{}

	dataInput := AddressUnblockedEventData{
		Reason: "",
	}

	_, data, err := PackAddressUnblockedEvent(
		accountInput,
		dataInput,
	)
	require.NoError(t, err)

	unpacked, err := UnpackAddressUnblockedEventData(data)
	require.NoError(t, err)
	require.Equal(t, dataInput, unpacked)
}

func BenchmarkBlockList(b *testing.B) {
	// Benchmark tests.
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			test.Bench(b, Module, extstate.NewTestStateDB(b))
		})
	}
}
