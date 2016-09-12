package compiler

// The parser structure holds the parser's internal state.
type parser struct {
	scanner Scanner

	pos Pos    // token position
	tok Token  // one token look-ahead
	lit string // token literal
}

func (p *parser) init(src []byte) {
	p.scanner.Init(src)
	p.next()
}

func (p *parser) Parse() []interface{} {
	var cmds []interface{}
	for {
		if p.tok != EOF {
			cmds = append(cmds, p.parseCommand())
			p.next()
		}
	}
}

func (p *parser) parseCommand() interface{} {
	if p.tok.IsOperator() {
		switch p.tok {
		case IN, OUT, INC, DEC: // unary
			cmd := UnaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = p.tok
				return cmd
			}
		case LD, ST: // load/store
			cmd := BinaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = p.tok
				p.next()
				if p.tok.IsRegister() {
					cmd.Y = p.tok
					return cmd
				}
			}
		case XOR, ADD, SUB, MUL, DIV: // binary
			cmd := BinaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = p.tok
				p.next()
				if p.tok.IsRegister() || p.tok.IsLiteral() {
					cmd.Y = p.tok
					return cmd
				}
			}
		case B, BX, BZ, BXZ, BN, BXN: // branch
			cmd := UnaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() || p.tok.IsLiteral() {
				cmd.X = p.tok
				return cmd
			}
		case HLT:
			return Command{Op: HLT}
		}
	}

	return nil
}

func (p *parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}
