package main

import (
	"fmt"
	"math/big"
	mathbits "math/bits"
)

// VM represents the virtual machine.
type VM struct {
	R      [8]*big.Int   // Integer registers (R0-R7)
	F      [8]*big.Float // Floating-point registers (F0-F7)
	A      [8]uint32     // Address registers (A0-A7)
	SR     byte          // Status Register
	PC     uint32        // Program Counter
	J      uint32        // Jump Return Register
	Memory [64 * 1024 * 1024]byte
}

// NewVM initializes a new virtual machine.
func NewVM() *VM {
	vm := &VM{}
	for i := range vm.R {
		vm.R[i] = new(big.Int)
	}
	for i := range vm.F {
		vm.F[i] = (new(big.Float)).SetPrec(256)
	}
	return vm
}

// SetFlag sets a specific flag in the Status Register.
func (vm *VM) SetFlag(flag int, value bool) {
	if value {
		vm.SR |= 1 << flag
	} else {
		vm.SR &^= 1 << flag
	}
}

// GetFlag checks if a specific flag is set.
func (vm *VM) GetFlag(flag int) bool {
	return (vm.SR>>flag)&1 == 1
}

// ======================
// Memory Access Instructions
// ======================

func (vm *VM) Load(rd int, ax int) {
	addr := vm.A[ax]
	if rd < 8 {
		vm.R[rd].SetBytes(vm.Memory[addr : addr+32])
	} else if rd < 16 {
		bitsVal := new(big.Int).SetBytes(vm.Memory[addr : addr+32])
		vm.F[rd-8].SetInt(bitsVal)
	} else {
		vm.A[rd-16] = uint32(vm.Memory[addr])<<24 |
			uint32(vm.Memory[addr+1])<<16 |
			uint32(vm.Memory[addr+2])<<8 |
			uint32(vm.Memory[addr+3])
	}
}

func (vm *VM) Store(rs int, ax int) {
	addr := vm.A[ax]
	if rs < 8 {
		copy(vm.Memory[addr:addr+32], vm.R[rs].Bytes())
	} else if rs < 16 {
		bitsVal := new(big.Int)
		vm.F[rs-8].Int(bitsVal)
		copy(vm.Memory[addr:addr+32], bitsVal.Bytes())
	} else {
		val := vm.A[rs-16]
		vm.Memory[addr] = byte(val >> 24)
		vm.Memory[addr+1] = byte(val >> 16)
		vm.Memory[addr+2] = byte(val >> 8)
		vm.Memory[addr+3] = byte(val)
	}
}

// ======================
// Arithmetic Instructions
// ======================

func (vm *VM) Add(rd, rs, rt int) {
	if rd < 8 {
		res := new(big.Int).Add(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
		vm.SetFlag(OF, res.BitLen() > 256)
	} else {
		res := new(big.Float).Add(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
	}
}

func (vm *VM) Sub(rd, rs, rt int) {
	if rd < 8 {
		res := new(big.Int).Sub(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
		vm.SetFlag(OF, res.BitLen() > 256)
	} else {
		res := new(big.Float).Sub(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
	}
}

func (vm *VM) Mul(rd, rs, rt int) {
	if rd < 8 {
		res := new(big.Int).Mul(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
		vm.SetFlag(OF, res.BitLen() > 256)
	} else {
		res := new(big.Float).Mul(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
	}
}

func (vm *VM) Div(rd, rs, rt int) {
	if rd < 8 {
		if vm.R[rt].Sign() == 0 {
			vm.SetFlag(DF, true)
			return
		}
		res := new(big.Int).Div(vm.R[rs], vm.R[rt])
		vm.R[rd].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
	} else {
		if vm.F[rt-8].Sign() == 0 {
			vm.SetFlag(DF, true)
			return
		}
		res := new(big.Float).Quo(vm.F[rs-8], vm.F[rt-8])
		vm.F[rd-8].Set(res)
		vm.SetFlag(ZF, res.Sign() == 0)
	}
}

// ======================
// Comparison Instruction
// ======================

func (vm *VM) Compare(rs, rt int) {
	if rs < 8 {
		// Integer comparison
		cmp := vm.R[rs].Cmp(vm.R[rt])
		vm.SetFlag(ZF, cmp == 0)
		vm.SetFlag(LT, cmp < 0)
		vm.SetFlag(GT, cmp > 0)
	} else {
		// Floating-point comparison
		cmp := vm.F[rs-8].Cmp(vm.F[rt-8])
		vm.SetFlag(ZF, cmp == 0)
		vm.SetFlag(LT, cmp < 0)
		vm.SetFlag(GT, cmp > 0)
	}
}

// ======================
// Type Conversion Instructions
// ======================

func (vm *VM) ITOF(fd int, rs int) {
	floatVal := new(big.Float).SetInt(vm.R[rs])
	fmt.Println(vm.F, fd)
	vm.F[fd].Set(floatVal)
}

func (vm *VM) FTOI(rd int, fs int) {
	intVal := new(big.Int)
	vm.F[fs].Int(intVal)
	vm.R[rd].Set(intVal)
}

// ======================
// Logical/Bitwise Instructions
// ======================

func (vm *VM) And(rd, rs, rt int) {
	res := new(big.Int).And(vm.R[rs], vm.R[rt])
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

func (vm *VM) Or(rd, rs, rt int) {
	res := new(big.Int).Or(vm.R[rs], vm.R[rt])
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

func (vm *VM) Xor(rd, rs, rt int) {
	res := new(big.Int).Xor(vm.R[rs], vm.R[rt])
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

func (vm *VM) Not(rd, rs int) {
	res := new(big.Int).Not(vm.R[rs])
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

func (vm *VM) Lsh(rd, rs, n int) {
	res := new(big.Int).Lsh(vm.R[rs], uint(n))
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

func (vm *VM) Rsh(rd, rs, n int) {
	res := new(big.Int).Rsh(vm.R[rs], uint(n))
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

func (vm *VM) Csh(rd, rs, n int) {
	bitsVal := vm.R[rs].Bits()
	size := len(bitsVal) * mathbits.UintSize // 使用重命名后的 mathbits 包
	n %= size                                // Normalize shift amount
	if n < 0 {
		n += size
	}

	// Split into high and low parts
	high := new(big.Int).Rsh(vm.R[rs], uint(size-n))
	low := new(big.Int).Lsh(vm.R[rs], uint(n))
	res := new(big.Int).Or(low, high)
	vm.R[rd].Set(res)
	vm.SetFlag(ZF, res.Sign() == 0)
}

// ======================
// Control Flow Instructions
// ======================

func (vm *VM) Jump(addr uint32) {
	vm.J = vm.PC + 1
	vm.PC = addr
}

func (vm *VM) JumpIf(addr uint32, condition bool) {
	if condition {
		vm.Jump(addr)
	}
}

// ======================
// Instruction Execution
// ======================

func (vm *VM) Execute(instruction uint32) {
	opcode := (instruction >> 24) & 0xFF
	rd := (instruction >> 20) & 0xF
	rs := (instruction >> 16) & 0xF
	rt := (instruction >> 12) & 0xF
	ax := (instruction >> 8) & 0xF
	imm := instruction & 0xFF

	switch opcode {
	case OP_NOP:
		// Do nothing
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
	case OP_JZ:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(ZF))
	case OP_JNZ:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), !vm.GetFlag(ZF))
	case OP_JGT:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(GT))
	case OP_JLT:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(LT))
	case OP_JEQ:
		vm.JumpIf(uint32(instruction&0x00FFFFFF), vm.GetFlag(ZF))
	default:
		fmt.Printf("Unknown opcode: %02X\n", opcode)
	}
	vm.PC++
}
