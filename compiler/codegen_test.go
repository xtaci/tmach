package compiler

import "testing"

func TestGenerate(t *testing.T) {
	p := Parser{}
	p.Init([]byte(code1))
	cmds := p.Parse()
	buf := Generate(cmds)
	t.Log(buf.Bytes())
}
