package tmach

import (
	"testing"

	"github.com/xtaci/tmach/compiler"
)

func TestVM(t *testing.T) {
	var src = `
		NOP
	L:
		IN R0
		OUT R0
		B L
	`
	p := compiler.Parser{}
	p.Init([]byte(src))
	cmds := p.Parse()
	buf := compiler.Generate(cmds)

	prog := newProgram(1024, 1024)
	prog.Load(buf.Bytes())
	go prog.Run()
	const N = 10
	go func() {
		for i := 0; i < N; i++ {
			prog.IN <- int32(i)
		}
	}()

	for i := 0; i < N; i++ {
		t.Log(<-prog.OUT)
	}
}