package compiler

import "testing"

func TestParser(t *testing.T) {
	var src = `
	L:
		IN R0
		IMUL R0 2
		OUT R0
		B L
	`
	p := Parser{}
	p.Init([]byte(src))
	cmds := p.Parse()
	t.Log(cmds)
}
