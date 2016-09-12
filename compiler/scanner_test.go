package compiler

import (
	"fmt"
	"testing"
)

var src = `
IN R0
OUT R0
LDR 0 R1
JMP R1
`

func TestScanner(t *testing.T) {
	s := Scanner{}
	s.Init([]byte(src))
	pos, tok, lit := s.Scan()
	for tok != EOF && tok != ILLEGAL {
		fmt.Println(pos, tok, lit)
		pos, tok, lit = s.Scan()
	}
}
