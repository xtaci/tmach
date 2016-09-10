package tmach

type Program struct {
	reg  [REGNUM]int32
	mem  []byte
	code []byte
	in   chan int32
	out  chan int32
}

func newProgram(memsize int, code []byte) *Program {
	p := new(Program)
	p.mem = make([]byte, memsize)
	p.code = code
	p.in = make(chan int32)
	p.out = make(chan int32)
	return p
}

func (p *Program) run() {
	pc := 0
	for p.code[pc] != HLT {
		opcode := p.code[pc]
		switch opcode {
		case IN:
		case OUT:
		}
	}
}
