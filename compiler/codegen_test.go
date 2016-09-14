package compiler

import "testing"

func TestGenerate(t *testing.T) {
	p := Parser{}
	p.Init([]byte(code2))
	cmds := p.Parse()
	buf := Generate(cmds)
	t.Log(buf.Bytes())
}
