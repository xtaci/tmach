package compiler

type Label struct {
	Name string
}

type Operation struct {
	Op       Token
	Operands []Operand
}

type Operand interface{}

type IntOperand struct {
	Value int64
}

type IdentOperand struct {
	Name string
}

type RegisterOperand struct {
	Name Token
}
