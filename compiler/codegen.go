package compiler

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/xtaci/tmach/arch"
)

var cpuCodes = [...]int{
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

	NOP: arch.NOP,
	IN:  arch.IN,
	OUT: arch.OUT,
	LD:  arch.LD,
	ST:  arch.ST,
	XOR: arch.XOR,
	ADD: arch.ADD,
	SUB: arch.SUB,
	MUL: arch.MUL,
	DIV: arch.DIV,
	JMP: arch.JMP,
	JN:  arch.JN,
	JR:  arch.JR,
	JRN: arch.JRN,
	HLT: arch.HLT,
}

func Generate(commands []interface{}) *bytes.Buffer {
	labels := make(map[string]uintptr)
	offset := uintptr(0)
	code := new(bytes.Buffer)

	for k := range commands {
		switch typedCmd := commands[k].(type) {
		case Label:
			labels[typedCmd.Name] = offset
		case Operation:
			switch typedCmd.Op {
			case IN, OUT: // IO operations
				opcode := (arch.TYPE_IO << arch.TypeShift)
				opcode |= cpuCodes[typedCmd.Op] << arch.OpShift
				opcode |= cpuCodes[typedCmd.Operands[0].(RegisterOperand).Name] << arch.RnShift
				binary.Write(code, binary.LittleEndian, uint16(opcode))
				offset += unsafe.Sizeof(opcode)
			case NOP, HLT:
				opcode := (arch.TYPE_IO << arch.TypeShift)
				opcode |= cpuCodes[typedCmd.Op] << arch.OpShift
				binary.Write(code, binary.LittleEndian, uint16(opcode))
				offset += unsafe.Sizeof(opcode)
			case LD, ST:
				opcode := (arch.TYPE_MEM << arch.TypeShift)
				opcode |= cpuCodes[typedCmd.Op] << arch.OpShift
				opcode |= cpuCodes[typedCmd.Operands[0].(RegisterOperand).Name] << arch.RnShift
				immop, ok := typedCmd.Operands[1].(IntOperand)
				if ok {
					opcode |= 1 << arch.ImmShift
				}
				binary.Write(code, binary.LittleEndian, uint16(opcode))
				offset += unsafe.Sizeof(opcode)
				if ok {
					binary.Write(code, binary.LittleEndian, immop.Value)
					offset += unsafe.Sizeof(immop.Value)
				}
			case XOR, ADD, SUB, MUL, DIV:
				opcode := (arch.TYPE_ALU << arch.TypeShift)
				opcode |= cpuCodes[typedCmd.Op] << arch.OpShift
				opcode |= cpuCodes[typedCmd.Operands[0].(RegisterOperand).Name] << arch.RnShift
				immop, ok := typedCmd.Operands[1].(IntOperand)
				if ok {
					opcode |= 1 << arch.ImmShift
				}
				binary.Write(code, binary.LittleEndian, uint16(opcode))
				offset += unsafe.Sizeof(opcode)
				if ok {
					binary.Write(code, binary.LittleEndian, immop.Value)
					offset += unsafe.Sizeof(immop.Value)
				}
			case JMP, JN:
				opcode := (arch.TYPE_BRANCH << arch.TypeShift)
				opcode |= cpuCodes[typedCmd.Op] << arch.OpShift
				opcode |= 1 << arch.ImmShift
				binary.Write(code, binary.LittleEndian, uint16(opcode))
				binary.Write(code, binary.LittleEndian, labels[typedCmd.Operands[0].(IdentOperand).Name])
				offset += 6
			case JR, JRN:
				opcode := (arch.TYPE_BRANCH << arch.TypeShift)
				opcode |= cpuCodes[typedCmd.Op] << arch.OpShift
				opcode |= cpuCodes[typedCmd.Operands[0].(RegisterOperand).Name] << arch.RnShift
				binary.Write(code, binary.LittleEndian, uint16(opcode))
				offset += unsafe.Sizeof(opcode)
			}
		}
	}
	return code
}
