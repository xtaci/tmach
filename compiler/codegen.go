package compiler

import (
	"bytes"
	"encoding/binary"

	"github.com/xtaci/tmach"
)

var cpuCodes = [...]byte{
	R0:  tmach.R0,
	R1:  tmach.R1,
	R2:  tmach.R2,
	R3:  tmach.R3,
	R4:  tmach.R4,
	R5:  tmach.R5,
	R6:  tmach.R6,
	R7:  tmach.R7,
	R8:  tmach.R8,
	R9:  tmach.R9,
	R10: tmach.R10,
	R11: tmach.R11,
	R12: tmach.R12,
	R13: tmach.R13,
	R14: tmach.R14,
	R15: tmach.R15,

	NOP:  tmach.NOP,
	IN:   tmach.IN,
	OUT:  tmach.OUT,
	LD:   tmach.LD,
	ST:   tmach.ST,
	XOR:  tmach.XOR,
	ADD:  tmach.ADD,
	SUB:  tmach.SUB,
	MUL:  tmach.MUL,
	DIV:  tmach.DIV,
	IXOR: tmach.IXOR,
	IADD: tmach.IADD,
	ISUB: tmach.ISUB,
	IMUL: tmach.IMUL,
	IDIV: tmach.IDIV,
	INC:  tmach.INC,
	DEC:  tmach.DEC,
	B:    tmach.B,
	BZ:   tmach.BZ,
	BN:   tmach.BN,
	BX:   tmach.BX,
	BXZ:  tmach.BXZ,
	BXN:  tmach.BXN,
	HLT:  tmach.HLT,
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
