package compiler

import "testing"

func TestParser(t *testing.T) {
	p := Parser{}
	p.Init([]byte(errCode1))
	cmds := p.Parse()
	t.Logf("%+v", cmds)
}
