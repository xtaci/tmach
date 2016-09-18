package compiler

import (
	"log"
	"strconv"

	"github.com/pkg/errors"
)

// The parser structure holds the parser's internal state.
type Parser struct {
	file    *File
	scanner Scanner
	pos     Pos    // token position
	tok     Token  // one token look-ahead
	lit     string // token literal
}

func (p *Parser) Init(src []byte) {
	fset := NewFileSet()
	p.file = fset.AddFile("", fset.Base(), len(src))
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
	switch p.tok {
	case IDENT:
		lit := p.lit
		p.next()
		if p.tok != COLON {
			p.errorExpected(p.pos, "'"+COLON.String()+"'")
		}
		return Label{lit}
	case IN, OUT:
		cmd := Operation{}
		cmd.Op = p.tok
		p.next()
		if p.tok.IsRegister() {
			cmd.Operands = append(cmd.Operands, RegisterOperand{Name: p.tok})
			return cmd
		}
	case LD, ST: // load/store
		cmd := Operation{}
		cmd.Op = p.tok
		p.next()
		if p.tok.IsRegister() {
			cmd.Operands = append(cmd.Operands, RegisterOperand{Name: p.tok})
			p.next()
			p.expect(COMMA)
			if p.tok.IsRegister() {
				cmd.Operands = append(cmd.Operands, RegisterOperand{Name: p.tok})
				return cmd
			} else if p.tok.IsLiteral() {
				i, _ := strconv.ParseInt(p.lit, 0, 32)
				cmd.Operands = append(cmd.Operands, IntOperand{Value: int64(i)})
				return cmd
			}
		}
	case XOR, ADD, SUB, MUL, DIV:
		cmd := Operation{}
		cmd.Op = p.tok
		t := false
		if p.tok == XOR {
			t = true
		}
		if t {
			log.Println(p.tok)
		}
		p.next()
		if t {
			log.Println(p.tok)
		}
		if p.tok.IsRegister() {
			cmd.Operands = append(cmd.Operands, RegisterOperand{Name: p.tok})
			if t {
				log.Println(p.tok)
			}
			p.next()
			if t {
				log.Println(p.tok)
			}
			p.expect(COMMA)
			if t {
				log.Println(p.tok)
			}
			if p.tok.IsRegister() {
				cmd.Operands = append(cmd.Operands, RegisterOperand{Name: p.tok})
				return cmd
			} else if p.tok.IsLiteral() {
				i, _ := strconv.ParseInt(p.lit, 0, 32)
				cmd.Operands = append(cmd.Operands, IntOperand{Value: int64(i)})
				return cmd
			}
		}
	case JMP, JN, JZ: // branch
		cmd := Operation{}
		cmd.Op = p.tok
		p.next()
		if p.tok.IsLiteral() {
			cmd.Operands = append(cmd.Operands, IdentOperand{p.lit})
			return cmd
		}
	case JR, JRN, JRZ:
		cmd := Operation{}
		cmd.Op = p.tok
		p.next()
		if p.tok.IsRegister() {
			cmd.Operands = append(cmd.Operands, RegisterOperand{Name: p.tok})
			return cmd
		}
	case NOP, HLT:
		return Operation{Op: p.tok}
	}

	return nil
}

func (p *Parser) expect(tok Token) Pos {
	pos := p.pos
	if p.tok != tok {
		p.errorExpected(pos, "'"+tok.String()+"'")
	}
	p.next() // make progress
	return pos
}

func (p *Parser) errorExpected(pos Pos, msg string) {
	msg = "expected " + msg
	if pos == p.pos {
		msg += ", found '" + p.tok.String() + "'"
		if p.tok.IsLiteral() {
			msg += " " + p.lit
		}
	}
	p.error(pos, msg)
}

type bailout struct{}

func (p *Parser) error(pos Pos, msg string) {
	log.Printf("%+v, %+v", pos, errors.New(msg))
	/*
		epos := p.file.Position(pos)

		// If AllErrors is not set, discard errors reported on the same line
		// as the last recorded error and stop parsing if there are more than
		// 10 errors.
		if p.mode&AllErrors == 0 {
			n := len(p.errors)
			if n > 0 && p.errors[n-1].Pos.Line == epos.Line {
				return // discard - likely a spurious error
			}
			if n > 10 {
				panic(bailout{})
			}
		}

		p.errors.Add(epos, msg)
	*/
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}
