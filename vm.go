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
			//log.Println(p.data)
			*pc++
			switch opcode {
			case arch.NOP:
			case arch.IN:
				Rn := p.code[*pc]
				*pc++
				p.reg[Rn] = <-p.IN
			case arch.OUT:
				Rn := p.code[*pc]
				*pc++
				p.OUT <- p.reg[Rn]
			case arch.B:
				off := int32(binary.LittleEndian.Uint32(p.code[*pc:]))
				*pc = off
			case arch.LD:
				Rn := p.code[*pc]
				*pc++
				Rm := p.code[*pc]
				*pc++
				p.reg[Rn] = p.data[p.reg[Rm]]
			case arch.ST:
				Rn := p.code[*pc]
				*pc++
				Rm := p.code[*pc]
				*pc++
				p.data[p.reg[Rm]] = p.reg[Rn]
			case arch.INC:
				Rn := p.code[*pc]
				*pc++
				p.reg[Rn]++
			case arch.DEC:
				Rn := p.code[*pc]
				*pc++
				p.reg[Rn]--
			case arch.XOR:
				Rn := p.code[*pc]
				*pc++
				Rm := p.code[*pc]
				*pc++
				p.reg[Rn] ^= p.reg[Rm]
			case arch.IMUL:
				reg := p.code[*pc]
				v := p.reg[reg]
				*pc++
				imm := int32(binary.LittleEndian.Uint32(p.code[*pc:]))
				*pc++
				p.reg[reg] = imm * v
			default:
				println("illegal", opcode)
			}
		}
	}
}
