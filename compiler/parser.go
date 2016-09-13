package compiler

// The parser structure holds the parser's internal state.
type Parser struct {
	scanner Scanner
	offset  int32  // instruction offset
	pos     Pos    // token position
	tok     Token  // one token look-ahead
	lit     string // token literal
	labels  map[string]int32
}

func (p *Parser) Init(src []byte) {
	p.scanner.Init(src)
	p.next()
}

func (p *Parser) Parse() []interface{} {
	var cmds []interface{}
	for {
		if p.tok != EOF {
			cmds = append(cmds, p.parseCommand())
			p.next()
		} else {
			return cmds
		}
	}
}

func (p *Parser) parseCommand() interface{} {
	if p.tok.IsOperator() {
		switch p.tok {
		case IDENT:
			lit := p.lit
			p.next()
			if p.tok == COLON {
				p.labels[lit] = p.offset
			}
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
				if p.tok.IsRegister() {
					cmd.Y = p.tok
					return cmd
				}
			}
		case IXOR, IADD, ISUB, IMUL, IDIV: // binary
			cmd := BinaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = p.tok
				p.next()
				if p.tok.IsLiteral() {
					cmd.Y = p.tok
					return cmd
				}
			}
		case B, BZ, BN: // branch
			cmd := UnaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsLiteral() {
				cmd.X = p.tok
				return cmd
			}
		case BX, BXZ, BXN:
			cmd := UnaryCommand{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = p.tok
				return cmd
			}
		case HLT:
			return Command{Op: HLT}
		}
	}

	return nil
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}
