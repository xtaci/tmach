package compiler

import "testing"

func TestGenerate(t *testing.T) {
	var src = `
	L:
		IN R0
		OUT R0
		B L
	`
	p := Parser{}
	p.Init([]byte(src))
	cmds := p.Parse()
	buf := Generate(cmds, p.labels)
	t.Log(buf.Bytes())
}
