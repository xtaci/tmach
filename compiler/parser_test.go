package compiler

import (
	"log"
	"testing"
)

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func TestParser(t *testing.T) {
	p := Parser{}
	p.Init([]byte(code1))
	cmds := p.Parse()
	t.Logf("%+v", cmds)
}
