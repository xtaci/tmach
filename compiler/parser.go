package compiler

import "strconv"

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
			if cmd := p.parseCommand(); cmd != nil {
				cmds = append(cmds, cmd)
			}
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
			cmd := UnaryOp{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = RegisterOperand{Name: p.tok}
				return cmd
			}
		case LD, ST: // load/store
			cmd := BinaryOp{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = RegisterOperand{Name: p.tok}
				p.next()
				if p.tok.IsRegister() {
					cmd.Y = RegisterOperand{Name: p.tok}
					return cmd
				}
			}
		case XOR, ADD, SUB, MUL, DIV: // binary
			cmd := BinaryOp{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = RegisterOperand{Name: p.tok}
				p.next()
				if p.tok.IsRegister() {
					cmd.Y = RegisterOperand{Name: p.tok}
					return cmd
				}
			}
		case IXOR, IADD, ISUB, IMUL, IDIV: // binary
			cmd := BinaryOp{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = RegisterOperand{Name: p.tok}
				p.next()
				if p.tok.IsLiteral() {
					i, _ := strconv.ParseInt(p.lit, 0, 32)
					cmd.Y = IntOperand{Value: int32(i)}
					return cmd
				}
			}
		case B, BZ, BN: // branch
			cmd := UnaryOp{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsLiteral() {
				cmd.X = IdentOperand{p.lit}
				return cmd
			}
		case BX, BXZ, BXN:
			cmd := UnaryOp{}
			cmd.Op = p.tok
			p.next()
			if p.tok.IsRegister() {
				cmd.X = RegisterOperand{Name: p.tok}
				return cmd
			}
		case HLT:
			return OpCode{Op: HLT}
		}
	}

	return nil
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}
