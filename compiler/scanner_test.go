package compiler

import (
	"fmt"
	"testing"
)

func TestScanner(t *testing.T) {
	var src = `
	IN R0
	OUT R0
	B -2
	`
	s := Scanner{}
	s.Init([]byte(src))
	pos, tok, lit := s.Scan()
	for tok != EOF && tok != ILLEGAL {
		fmt.Println(pos, tok, lit)
		pos, tok, lit = s.Scan()
	}
}
