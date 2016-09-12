package compiler

type Command struct {
	Op Token
}

type UnaryCommand struct {
	Command
	X Token
}

type BinaryCommand struct {
	Command
	X Token
	Y Token
}
