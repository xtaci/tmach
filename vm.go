package tmach

import (
	"encoding/binary"

	"github.com/xtaci/tmach/arch"
)

type Program struct {
	reg  [arch.REGCOUNT]int32
	data []int32
	code []byte
	IN   chan int32
	OUT  chan int32
}

func newProgram(datasize, codesize int) *Program {
	p := new(Program)
	p.data = make([]int32, datasize)
	p.code = make([]byte, codesize)
	p.IN = make(chan int32)
	p.OUT = make(chan int32)
	return p
}

func (p *Program) Load(code []byte) {
	copy(p.code, code)
}

func (p *Program) Run() {
	pc := &p.reg[arch.PC]
	for {
		if *pc < int32(len(p.code)) && p.code[*pc] != arch.HLT {
			opcode := p.code[*pc]
			*pc++
			println("opcode:", opcode, *pc)
			switch opcode {
			case arch.NOP:
			case arch.IN:
				n := p.code[*pc]
				*pc++
				p.reg[n] = <-p.IN
			case arch.OUT:
				n := p.code[*pc]
				*pc++
				p.OUT <- p.reg[n]
			case arch.B:
				off := int32(binary.LittleEndian.Uint32(p.code[*pc:]))
				*pc = off
			case arch.LD:
				n := p.code[*pc]
				*pc++
				m := p.code[*pc]
				*pc++
				p.reg[m] = p.data[n]
			case arch.ST:
				n := p.code[*pc]
				*pc++
				m := p.code[*pc]
				*pc++
				p.data[m] = p.reg[n]
			default:
				println("illegal", opcode)
			}
		}
	}
}
