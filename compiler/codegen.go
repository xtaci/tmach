package compiler

import (
	"bytes"
	"encoding/binary"

	"github.com/xtaci/tmach/arch"
)

var cpuCodes = [...]byte{
	R0:  arch.R0,
	R1:  arch.R1,
	R2:  arch.R2,
	R3:  arch.R3,
	R4:  arch.R4,
	R5:  arch.R5,
	R6:  arch.R6,
	R7:  arch.R7,
	R8:  arch.R8,
	R9:  arch.R9,
	R10: arch.R10,
	R11: arch.R11,
	R12: arch.R12,
	R13: arch.R13,
	R14: arch.R14,
	R15: arch.R15,

	NOP:  arch.NOP,
	IN:   arch.IN,
	OUT:  arch.OUT,
	LD:   arch.LD,
	ST:   arch.ST,
	XOR:  arch.XOR,
	ADD:  arch.ADD,
	SUB:  arch.SUB,
	MUL:  arch.MUL,
	DIV:  arch.DIV,
	IXOR: arch.IXOR,
	IADD: arch.IADD,
	ISUB: arch.ISUB,
	IMUL: arch.IMUL,
	IDIV: arch.IDIV,
	INC:  arch.INC,
	DEC:  arch.DEC,
	B:    arch.B,
	BZ:   arch.BZ,
	BN:   arch.BN,
	BX:   arch.BX,
	BXZ:  arch.BXZ,
	BXN:  arch.BXN,
	HLT:  arch.HLT,
}

func Generate(commands []interface{}) *bytes.Buffer {
	labels := make(map[string]int32)
	offset := int32(0)
	code := new(bytes.Buffer)

	for k := range commands {
		switch typedCmd := commands[k].(type) {
		case Label:
			labels[typedCmd.Name] = offset
		case OpCode:
			code.WriteByte(cpuCodes[typedCmd.Op])
			offset++
		case UnaryOp:
			code.WriteByte(cpuCodes[typedCmd.Op])
			offset++
			switch typedOperand := typedCmd.X.(type) {
			case IdentOperand:
				binary.Write(code, binary.LittleEndian, labels[typedOperand.Name])
				offset += 4
			case IntOperand:
				binary.Write(code, binary.LittleEndian, typedOperand.Value)
				offset += 4
			case RegisterOperand:
				code.WriteByte(cpuCodes[typedOperand.Name])
				offset++
			}
		case BinaryOp:
			code.WriteByte(cpuCodes[typedCmd.Op])
			switch typedOperand := typedCmd.X.(type) {
			case IdentOperand:
				binary.Write(code, binary.LittleEndian, labels[typedOperand.Name])
				offset += 4
			case IntOperand:
				binary.Write(code, binary.LittleEndian, typedOperand.Value)
				offset += 4
			case RegisterOperand:
				code.WriteByte(cpuCodes[typedOperand.Name])
				offset++
			}
			switch typedOperand := typedCmd.Y.(type) {
			case IdentOperand:
				binary.Write(code, binary.LittleEndian, labels[typedOperand.Name])
				offset += 4
			case IntOperand:
				binary.Write(code, binary.LittleEndian, typedOperand.Value)
				offset += 4
			case RegisterOperand:
				code.WriteByte(cpuCodes[typedOperand.Name])
				offset++
			}
		}
	}
	return code
}
