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

func TestDouble(t *testing.T) {
	var src = `
		NOP
	L:
		IN R0
		IMUL R0 2
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

func TestReverse2(t *testing.T) {
	var src = `
		NOP
	L1:
		IN R0
		XOR R1 R1
		ST R0 R1
	L2:	
		IN R0
		INC R1
		ST R0 R1

	OUT:
		LD R0 R1
		OUT R0
		DEC R1	
		LD R0 R1
		OUT R0
		B L1
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
