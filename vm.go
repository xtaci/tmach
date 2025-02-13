package tmach

import (
	"fmt"
	"math/big"
)

// VM represents the tmach virtual machine.
type VM struct {
	// 256-bit general-purpose registers.
	// When used for integer operations, refer to them as R0-R7.
	// When used for floating-point operations, refer to them as F0-F7.
	R [8]*big.Int   // Integer registers R0-R7
	F [8]*big.Float // Floating-point registers F0-F7

	// 32-bit address registers (A0-A7) used for memory addressing.
	A [8]uint32

	// 8-bit status register.
	// Bit0: Zero Flag (ZF)
	// Bit1: Overflow Flag (OF)
	// Bit2: Divide-by-Zero Flag (DF)
	// Bits3-5: Comparison Result (CR) (e.g., LT, GT, EQ)
	// Bits6-7: Reserved
	SR byte

	// Program Counter (PC) holds the current instruction address.
	PC uint32

	// Jump Return Register (J) holds the return address after a jump.
	J uint32

	// Memory of the virtual machine (e.g., 64 MB).
	Memory [64 * 1024 * 1024]byte
}

// NewVM initializes and returns a new virtual machine.
func NewVM() *VM {
	vm := &VM{}
	for i := range vm.R {
		vm.R[i] = new(big.Int)
	}
	for i := range vm.F {
		// Set floating-point precision to 256 bits.
		vm.F[i] = new(big.Float).SetPrec(256)
	}
	return vm
}

// SetFlag sets or clears a specific flag in the status register.
func (vm *VM) SetFlag(flag int, value bool) {
	if value {
		vm.SR |= 1 << flag
	} else {
		vm.SR &^= 1 << flag
	}
}

// GetFlag returns true if the specified flag is set.
func (vm *VM) GetFlag(flag int) bool {
	return (vm.SR>>flag)&1 == 1
}

// ===================================================================
// Memory Access Instructions (Address is read directly from an address register)
// ===================================================================

// Load loads 32 bytes (256 bits) from memory at the address given by A[ax] into the target register.
// For integer registers (rd in 0..7), load into R; for floating-point registers (rd in 8..15), load into F.
func (vm *VM) Load(rd int, ax int) {
	addr := vm.A[ax]
	if addr+32 > uint32(len(vm.Memory)) {
		fmt.Println("Memory read out of bounds")
		return
	}
	if rd < 8 {
		// Load 256-bit integer value.
		vm.R[rd].SetBytes(vm.Memory[addr : addr+32])
	} else if rd < 16 {
		// Load 256-bit floating-point value.
		bitsVal := new(big.Int).SetBytes(vm.Memory[addr : addr+32])
		vm.F[rd-8].SetInt(bitsVal)
	} else {
		// If needed, address registers could be loaded here.
		// For now, we assume only R and F are used.
	}
}

// Store stores 32 bytes (256 bits) from the source register into memory at the address given by A[ax].
// For integer registers (rs in 0..7), store from R; for floating-point registers (rs in 8..15), store from F.
func (vm *VM) Store(rs int, ax int) {
	addr := vm.A[ax]
	if addr+32 > uint32(len(vm.Memory)) {
		fmt.Println("Memory write out of bounds")
		return
	}
	if rs < 8 {
		data := vm.R[rs].Bytes()
		// Pad to 32 bytes if necessary.
		padded := make([]byte, 32)
		copy(padded[32-len(data):], data)
		copy(vm.Memory[addr:addr+32], padded)
	} else if rs < 16 {
		bitsVal := new(big.Int)
		vm.F[rs-8].Int(bitsVal)
		data := bitsVal.Bytes()
		padded := make([]byte, 32)
		copy(padded[32-len(data):], data)
		copy(vm.Memory[addr:addr+32], padded)
	} else {
		// Not used in this design.
	}
}

// ===================================================================
// Arithmetic Instructions
// ===================================================================

// Add performs 256-bit addition. It uses integer registers (R) if rd < 8; otherwise, it uses floating-point registers.
func (vm *VM) Add(rd, rs, rt int) {
	if rd < 8 {
		res := new(big.Int).Add(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(0, res.Sign() == 0)    // Zero Flag
		vm.SetFlag(1, res.BitLen() > 256) // Overflow Flag
	} else {
		res := new(big.Float).Add(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(0, res.Sign() == 0) // Zero Flag
	}
}

// Sub performs subtraction.
func (vm *VM) Sub(rd, rs, rt int) {
	if rd < 8 {
		res := new(big.Int).Sub(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(0, res.Sign() == 0)
		vm.SetFlag(1, res.BitLen() > 256)
	} else {
		res := new(big.Float).Sub(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(0, res.Sign() == 0)
	}
}

// Mul performs multiplication.
func (vm *VM) Mul(rd, rs, rt int) {
	if rd < 8 {
		res := new(big.Int).Mul(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(0, res.Sign() == 0)
		vm.SetFlag(1, res.BitLen() > 256)
	} else {
		res := new(big.Float).Mul(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(0, res.Sign() == 0)
	}
}

// Div performs division. It checks for division by zero.
func (vm *VM) Div(rd, rs, rt int) {
	if rd < 8 {
		if vm.R[rt].Sign() == 0 {
			vm.SetFlag(2, true) // Divide-by-Zero Flag
			return
		}
		res := new(big.Int).Div(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(0, res.Sign() == 0)
	} else {
		// For floating-point, compare against 0.
		if vm.F[rt-8].Cmp(big.NewFloat(0)) == 0 {
			vm.SetFlag(2, true)
			return
		}
		res := new(big.Float).Quo(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(0, res.Sign() == 0)
	}
}

// ===================================================================
// Comparison Instruction
// ===================================================================

// Compare compares the values in Rs and Rt and sets the status flags accordingly.
// For integer registers, it updates Zero Flag, Less-Than (bit 3), and Greater-Than (bit 4) flags.
// For floating-point, similar behavior is applied.
func (vm *VM) Compare(rs, rt int) {
	if rs < 8 {
		cmp := vm.R[rs].Cmp(vm.R[rt])
		vm.SetFlag(0, cmp == 0) // Zero flag
		vm.SetFlag(3, cmp < 0)  // LT flag (bit 3)
		vm.SetFlag(4, cmp > 0)  // GT flag (bit 4)
	} else {
		cmp := vm.F[rs-8].Cmp(vm.F[rt-8])
		vm.SetFlag(0, cmp == 0)
		vm.SetFlag(3, cmp < 0)
		vm.SetFlag(4, cmp > 0)
	}
}

// ===================================================================
// Type Conversion Instructions
// ===================================================================

// ITOF converts an integer in register R[rs] to a floating-point number and stores it in F[fd].
func (vm *VM) ITOF(fd int, rs int) {
	floatVal := new(big.Float).SetInt(vm.R[rs])
	vm.F[fd].Set(floatVal)
}

// FTOI converts a floating-point number in register F[fs] to an integer and stores it in R[rd].
func (vm *VM) FTOI(rd int, fs int) {
	intVal := new(big.Int)
	vm.F[fs].Int(intVal)
	vm.R[rd].Set(intVal)
}

// ===================================================================
// Logical/Bitwise Instructions
// ===================================================================

// And performs a bitwise AND on R[rs] and R[rt] and stores the result in R[rd].
func (vm *VM) And(rd, rs, rt int) {
	res := new(big.Int).And(vm.R[rs], vm.R[rt])
	vm.R[rd].Set(res)
	vm.SetFlag(0, res.Sign() == 0)
}

// Or performs a bitwise OR on R[rs] and R[rt] and stores the result in R[rd].
func (vm *VM) Or(rd, rs, rt int) {
	res := new(big.Int).Or(vm.R[rs], vm.R[rt])
	vm.R[rd].Set(res)
	vm.SetFlag(0, res.Sign() == 0)
}

// Xor performs a bitwise XOR on R[rs] and R[rt] and stores the result in R[rd].
func (vm *VM) Xor(rd, rs, rt int) {
	res := new(big.Int).Xor(vm.R[rs], vm.R[rt])
	vm.R[rd].Set(res)
	vm.SetFlag(0, res.Sign() == 0)
}

// Not performs a bitwise NOT on R[rs] and stores the result in R[rd].
func (vm *VM) Not(rd, rs int) {
	// Create a mask for 256 bits
	mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

	// Perform bitwise NOT using XOR with the mask
	res := new(big.Int).Xor(vm.R[rs], mask)

	// Store the result in the destination register
	vm.R[rd].Set(res)

	// Set Zero Flag (ZF) if the result is zero
	vm.SetFlag(ZF, res.Sign() == 0)
}

// Lsh performs a logical left shift on R[rs] by n bits and stores the result in R[rd].
func (vm *VM) Lsh(rd, rs, n int) {
	res := new(big.Int).Lsh(vm.R[rs], uint(n))
	vm.R[rd].Set(res)
	vm.SetFlag(0, res.Sign() == 0)
}

// Rsh performs a logical right shift on R[rs] by n bits and stores the result in R[rd].
func (vm *VM) Rsh(rd, rs, n int) {
	res := new(big.Int).Rsh(vm.R[rs], uint(n))
	vm.R[rd].Set(res)
	vm.SetFlag(0, res.Sign() == 0)
}

// Csh performs a cyclic (rotational) shift on R[rs] by n bits and stores the result in R[rd].
// The immediate value n is treated as signed: positive for left rotation, negative for right rotation.
func (vm *VM) Csh(rd, rs, n int) {
	// Ensure the register value is represented in exactly 256 bits.
	data := vm.R[rs].Bytes()
	if len(data) < 32 {
		padded := make([]byte, 32)
		copy(padded[32-len(data):], data)
		data = padded
	}
	orig := new(big.Int).SetBytes(data)

	// Normalize shift amount (n mod 256)
	n = n % 256
	if n < 0 {
		n += 256
	}

	// Perform cyclic left rotation by n bits:
	// result = ((orig << n) OR (orig >> (256 - n))) mod (1 << 256)
	left := new(big.Int).Lsh(orig, uint(n))
	right := new(big.Int).Rsh(orig, uint(256-n))
	res := new(big.Int).Or(left, right)
	// Mask to 256 bits.
	mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	res.And(res, mask)
	vm.R[rd].Set(res)
	vm.SetFlag(0, res.Sign() == 0)
}

// ===================================================================
// Control Flow Instructions
// ===================================================================

// Jump performs an unconditional jump to the given address.
// It saves the return address (the next instruction) in the jump return register J.
func (vm *VM) Jump(addr uint32) {
	vm.J = vm.PC + 1 // Save the next instruction address in J
	vm.PC = addr     // Set PC to the target address
}

// JumpIf performs a jump to addr if condition is true.
func (vm *VM) JumpIf(addr uint32, condition bool) {
	if condition {
		vm.Jump(addr)
	}
}

// ===================================================================
// Instruction Execution
// ===================================================================

// Execute decodes and executes a 24-bit instruction stored in a uint32.
// The instruction is assumed to be in the following fixed format:
// [Opcode (8 bits)] [Field1 (4 bits)] [Field2 (4 bits)] [Field3 (4 bits)] [Field4 (4 bits)]
// For memory instructions, Field4 is reserved (set to 0).
// For jump instructions, the lower 24 bits represent the target address.
func (vm *VM) Execute(instruction uint32) {
	opcode := (instruction >> 24) & 0xFF
	rd := (instruction >> 20) & 0xF
	rs := (instruction >> 16) & 0xF
	rt := (instruction >> 12) & 0xF
	ax := (instruction >> 8) & 0xF
	imm := instruction & 0xFF

	jumped := false
	switch opcode {
	case OP_NOP:
		// No operation.
	case OP_LOAD:
		vm.Load(int(rd), int(ax))
	case OP_STORE:
		vm.Store(int(rs), int(ax))
	case OP_ADD:
		vm.Add(int(rd), int(rs), int(rt))
	case OP_SUB:
		vm.Sub(int(rd), int(rs), int(rt))
	case OP_MUL:
		vm.Mul(int(rd), int(rs), int(rt))
	case OP_DIV:
		vm.Div(int(rd), int(rs), int(rt))
	case OP_CMP:
		vm.Compare(int(rs), int(rt))
	case OP_ITOF:
		vm.ITOF(int(rd), int(rs))
	case OP_FTOI:
		vm.FTOI(int(rd), int(rs))
	case OP_AND:
		vm.And(int(rd), int(rs), int(rt))
	case OP_OR:
		vm.Or(int(rd), int(rs), int(rt))
	case OP_XOR:
		vm.Xor(int(rd), int(rs), int(rt))
	case OP_NOT:
		vm.Not(int(rd), int(rs))
	case OP_LSH:
		vm.Lsh(int(rd), int(rs), int(imm))
	case OP_RSH:
		vm.Rsh(int(rd), int(rs), int(imm))
	case OP_CSH:
		vm.Csh(int(rd), int(rs), int(imm))
	case OP_JMP:
		vm.Jump(uint32(instruction & 0x00FFFFFF))
		jumped = true
	case OP_JZ:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(ZF))
		jumped = true
	case OP_JNZ:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), !vm.GetFlag(ZF))
		jumped = true
	case OP_JGT:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(GT))
		jumped = true
	case OP_JLT:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(LT))
		jumped = true
	case OP_JEQ:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(ZF))
		jumped = true
	default:
		fmt.Printf("Unknown opcode: %02X\n", opcode)
	}
	if !jumped {
		vm.PC++
	}
}
