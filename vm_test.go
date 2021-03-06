package tmach

import (
	"testing"

	"github.com/xtaci/tmach/compiler"
)

func TestCopy(t *testing.T) {
	var src = `
		NOP
	L:
		IN R0
		OUT R0
		JMP L
	`
	p := compiler.Parser{}
	p.Init([]byte(src))
	cmds := p.Parse()
	buf := compiler.Generate(cmds)

	prog := newProgram(1024, 1024)
	prog.Load(buf.Bytes())
	t.Log(buf.Bytes())
	go prog.Run()
	const N = 10
	go func() {
		for i := int64(0); i < N; i++ {
			prog.IN <- i
		}
	}()

	for i := 0; i < N; i++ {
		t.Log(<-prog.OUT)
	}
}

func TestDouble(t *testing.T) {
	var src = `
		NOP
	L:
		IN R0
		MUL R0,2
		OUT R0
		JMP L
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
		for i := int64(0); i < N; i++ {
			prog.IN <- i
		}
	}()

	for i := 0; i < N; i++ {
		t.Log(<-prog.OUT)
	}
}

func TestReverse2(t *testing.T) {
	var src = `
		NOP
	L1:
		IN R0
		ST R0,0
	L2:	
		IN R0
		ST R0,4

	OUT:
		LD R0,4
		OUT R0
		LD R0,0
		OUT R0
		JMP L1
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
		for i := int64(0); i < N; i++ {
			prog.IN <- i
		}
	}()

	for i := 0; i < N; i++ {
		t.Log(<-prog.OUT)
	}
}
