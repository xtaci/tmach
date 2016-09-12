package compiler

import "testing"

func TestParser(t *testing.T) {
	var src = `
	IN R0
	OUT R0
	B -2
	`
	p := Parser{}
	p.Init([]byte(src))
	cmds := p.Parse()
	t.Log(cmds)
}
