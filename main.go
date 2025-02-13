package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Create a new instance of the virtual machine
	vm := NewVM()

	// Initialize Address Register A0 with the memory address 0x1000 (4096 in decimal)
	vm.A[0] = 0x1000

	// Prepare a 256-bit integer value representing the number 1.
	// We use big.Int to create a high-precision integer.
	initVal := new(big.Int).SetInt64(1)
	valBytes := initVal.Bytes()

	// Pad the value to ensure it occupies 32 bytes (256 bits)
	padded := make([]byte, 32)
	copy(padded[32-len(valBytes):], valBytes)

	// Write the 32-byte value into memory at the address specified by A0
	addr := vm.A[0]
	copy(vm.Memory[addr:addr+32], padded)

	// Build a simple program as a slice of 32-bit instructions.
	// The program does the following:
	// 1. LOAD R0, [A0]   : Load the 256-bit value from memory (address in A0) into register R0.
	// 2. ADD R0, R0, R0   : Add R0 to itself (doubling the value in R0).
	// 3. STORE R0, [A0]   : Store the result from R0 back into memory at the address in A0.
	// 4. NOP              : No Operation.
	//
	// Each instruction is encoded in 24 bits (3 bytes). For memory instructions, no offset is used;
	// the effective address is determined entirely by the address register.
	program := []uint32{
		// Encoding for "LOAD R0, [A0]":
		// Opcode = OP_LOAD, Target Register (R0) = 0, Address Register (A0) = 0, Reserved = 0.
		(OP_LOAD << 24) | (0 << 20) | (0 << 8),
		// Encoding for "ADD R0, R0, R0":
		// Opcode = OP_ADD, Source Register 1 (R0) = 0, Source Register 2 (R0) = 0, Destination Register (R0) = 0.
		(OP_ADD << 24) | (0 << 20) | (0 << 16) | (0 << 12),
		// Encoding for "STORE R0, [A0]":
		// Opcode = OP_STORE, Source Register (R0) = 0, Address Register (A0) = 0, Reserved = 0.
		(OP_STORE << 24) | (0 << 16) | (0 << 8),
		// Encoding for "NOP":
		// Opcode = OP_NOP, remaining bits = 0.
		(OP_NOP << 24),
	}

	// Execute the program instruction by instruction.
	fmt.Println("Starting program execution...")
	for vm.PC < uint32(len(program)) {
		inst := program[vm.PC]
		fmt.Printf("PC: %d, Instruction: 0x%08X\n", vm.PC, inst)
		vm.Execute(inst)
	}

	// After executing the program, read the 32-byte value from memory at address A0
	// and convert it to a big.Int for display.
	result := new(big.Int).SetBytes(vm.Memory[vm.A[0] : vm.A[0]+32])
	fmt.Printf("Result stored at memory address 0x%X: %s\n", vm.A[0], result.String())
}
