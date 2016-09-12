package tmach

type Program struct {
	reg  [REGCOUNT]int32
	data []int32
	code []byte
	in   chan int32
	out  chan int32
}

func newProgram(datasize, codesize int) *Program {
	p := new(Program)
	p.data = make([]int32, datasize)
	p.code = make([]byte, codesize)
	p.in = make(chan int32)
	p.out = make(chan int32)
	return p
}

func (p *Program) Load(code []byte) {
	copy(p.code, code)
}

func (p *Program) run() {
	pc := 0
	for p.code[pc] != HLT {
		opcode := p.code[pc]
		switch opcode {
		case IN:
			n := p.code[pc+1]
			p.reg[n] = <-p.in
			pc += 2
		case OUT:
			n := p.code[pc+1]
			p.in <- p.reg[n]
			pc += 2
		case LD:
			n := p.code[pc+1]
			m := p.code[pc+2]
			p.reg[m] = p.data[n]
			pc += 3
		case ST:
			n := p.code[pc+1]
			m := p.code[pc+2]
			p.data[m] = p.reg[n]
			pc += 3
		}
	}
}
