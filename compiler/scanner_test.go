package compiler

import "testing"

func TestScanner(t *testing.T) {
	s := Scanner{}
	s.Init([]byte(code1))
	pos, tok, lit := s.Scan()
	for tok != EOF && tok != ILLEGAL {
		t.Logf("%+v, %+v, %+v", pos, tok, lit)
		pos, tok, lit = s.Scan()
	}
}
