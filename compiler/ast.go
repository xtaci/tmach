package compiler

type Label struct {
	Name string
}

type OpCode struct {
	Op Token
}

type Operand interface {
}

type IntOperand struct {
	Value int32
}

type IdentOperand struct {
	Name string
}

type RegisterOperand struct {
	Name Token
}

type UnaryOp struct {
	OpCode
	X Operand
}

type BinaryOp struct {
	OpCode
	X Operand
	Y Operand
}
