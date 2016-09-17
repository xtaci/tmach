package tmach

import (
	"encoding/binary"

	"github.com/xtaci/tmach/arch"
)

type Program struct {
	reg  [arch.REGISTER_COUNT]uint32
	data []byte
	code []byte
	IN   chan uint32
	OUT  chan uint32
}

func newProgram(datasize, codesize int) *Program {
	p := new(Program)
	p.data = make([]byte, datasize)
	p.code = make([]byte, codesize)
	p.IN = make(chan uint32)
	p.OUT = make(chan uint32)
	return p
}

func (p *Program) Load(code []byte) {
	copy(p.code, code)
}

func (p *Program) Run() {
	pc := &p.reg[arch.PC]
	for {
		if *pc < uint32(len(p.code)) && p.code[*pc] != arch.HLT {
			opcode := binary.LittleEndian.Uint16(p.code[*pc:])
			switch (opcode & arch.TypeMask) >> arch.TypeShift {
			case arch.TYPE_IO:
				p.execIO(opcode)
			case arch.TYPE_ALU:
				p.execALU(opcode)
			case arch.TYPE_MEM:
				p.execMEM(opcode)
			case arch.TYPE_BRANCH:
				p.execBranch(opcode)
			}
		}
	}
}

func (p *Program) execIO(opcode uint16) {
	pc := &p.reg[arch.PC]
	switch (opcode & arch.OpMask) >> arch.OpShift {
	case arch.NOP:
		*pc += 2
	case arch.IN:
		*pc += 2
		Rn := opcode & arch.RnMask
		p.reg[Rn] = <-p.IN
	case arch.OUT:
		*pc += 2
		Rn := opcode & arch.RnMask
		p.OUT <- p.reg[Rn]
	}
}

func (p *Program) execBranch(opcode uint16) {
	pc := &p.reg[arch.PC]
	Rn := (opcode & arch.RnMask) >> arch.RnShift
	switch (opcode & arch.OpMask) >> arch.OpShift {
	case arch.B:
		*pc += 2
		*pc = binary.LittleEndian.Uint32(p.code[*pc:])
	case arch.BZ:
		*pc += 2
		if p.reg[arch.PSR]&arch.PSR_ZERO == arch.PSR_ZERO {
			*pc = binary.LittleEndian.Uint32(p.code[*pc:])
		}
	case arch.BN:
		*pc += 2
		if p.reg[arch.PSR]&arch.PSR_NEG == arch.PSR_NEG {
			*pc = binary.LittleEndian.Uint32(p.code[*pc:])
		}
	case arch.BX:
		*pc = p.reg[Rn]
	case arch.BXZ:
		if p.reg[arch.PSR]&arch.PSR_ZERO == arch.PSR_ZERO {
			*pc = p.reg[Rn]
		} else {
			*pc += 2
		}
	case arch.BXN:
		if p.reg[arch.PSR]&arch.PSR_NEG == arch.PSR_NEG {
			*pc = p.reg[Rn]
		} else {
			*pc += 2
		}
	}
}

func (p *Program) execALU(opcode uint16) {
	pc := &p.reg[arch.PC]
	Rn := (opcode & arch.RnMask) >> arch.RnShift
	Rm := opcode & arch.RmMask
	Imm := (opcode & arch.ImmMask) >> arch.ImmShift
	switch (opcode & arch.OpMask) >> arch.OpShift {
	case arch.XOR:
		*pc += 2
		if Imm == 0 {
			p.reg[Rn] ^= p.reg[Rm]
		} else {
			p.reg[Rn] ^= binary.LittleEndian.Uint32(p.code[*pc:])
			*pc += 4
		}
	case arch.MUL:
		*pc += 2
		if Imm == 0 {
			p.reg[Rn] *= p.reg[Rm]
		} else {
			p.reg[Rn] *= binary.LittleEndian.Uint32(p.code[*pc:])
			*pc += 4
		}
	case arch.DIV:
		*pc += 2
		if Imm == 0 {
			p.reg[Rn] /= p.reg[Rm]
		} else {
			p.reg[Rn] /= binary.LittleEndian.Uint32(p.code[*pc:])
			*pc += 4
		}
	case arch.ADD:
		*pc += 2
		if Imm == 0 {
			p.reg[Rn] += p.reg[Rm]
		} else {
			p.reg[Rn] += binary.LittleEndian.Uint32(p.code[*pc:])
			*pc += 4
		}
	case arch.SUB:
		*pc += 2
		if Imm == 0 {
			p.reg[Rn] -= p.reg[Rm]
		} else {
			p.reg[Rn] -= binary.LittleEndian.Uint32(p.code[*pc:])
			*pc += 4
		}
	}
}

func (p *Program) execMEM(opcode uint16) {
	pc := &p.reg[arch.PC]
	Rn := (opcode & arch.RnMask) >> arch.RnShift
	Rm := opcode & arch.RmMask
	Imm := (opcode & arch.ImmMask) >> arch.ImmShift

	switch (opcode & arch.OpMask) >> arch.OpShift {
	case arch.LD:
		*pc += 2
		if Imm == 0 {
			addr := p.reg[Rm]
			p.reg[Rn] = binary.LittleEndian.Uint32(p.data[addr:])
		} else {
			addr := binary.LittleEndian.Uint32(p.code[*pc:])
			p.reg[Rn] = binary.LittleEndian.Uint32(p.data[addr:])
			*pc += 4
		}
	case arch.ST:
		*pc += 2
		if Imm == 0 {
			addr := p.reg[Rm]
			binary.LittleEndian.PutUint32(p.data[addr:], p.reg[Rn])
		} else {
			addr := binary.LittleEndian.Uint32(p.code[*pc:])
			binary.LittleEndian.PutUint32(p.data[addr:], p.reg[Rn])
			*pc += 4
		}
	}
}
