package compiler

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	INT
	FLOAT
	STRING
	literal_end

	reg_beg
	R0
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	reg_end

	opcode_beg
	IN  // IN REGn
	OUT // OUT REGn
	LD  // LD [REGn], REGm
	ST  // ST REGn, [REGm]
	LDR // LDR IMM, REGn
	XOR // XOR REGn, REGm
	ADD // ADD REGn, REGm
	SUB // SUB REGn, REGm
	MUL // MUL REGn, REGm
	DIV // DIV REGn, REGm
	INC // INC REGn
	DEC // DEC REGn
	JMP // JMP IMM
	JZ  // JZ IMM
	JGZ // JGZ IMM
	JLZ // JLZ IMM
	HLT
	opcode_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	R0:  "R0",
	R1:  "R1",
	R2:  "R2",
	R3:  "R3",
	R4:  "R4",
	R5:  "R5",
	R6:  "R6",
	R7:  "R7",
	R8:  "R8",
	R9:  "R9",
	R10: "R10",
	R11: "R11",
	R12: "R12",
	R13: "R13",
	R14: "R14",
	R15: "R15",

	IN:  "IN",
	OUT: "OUT",
	LD:  "LD",
	ST:  "ST",
	LDR: "LDR",
	XOR: "XOR",
	ADD: "AND",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",
	INC: "INC",
	DEC: "DEC",
	JMP: "JMP",
	JZ:  "JZ",
	JGZ: "JGZ",
	JLZ: "JLZ",
	HLT: "HLT",
}

type Token int

func (t *Token) IsOperator() bool {
	return false
}

func (t *Token) IsLiteral() bool {
	return false
}

func (t *Token) String() string {
	return ""
}

var opcodes map[string]Token
var registers map[string]Token

func init() {
	opcodes = make(map[string]Token)
	for i := opcode_beg + 1; i < opcode_end; i++ {
		opcodes[tokens[i]] = i
	}
	registers = make(map[string]Token)
	for i := reg_beg + 1; i < reg_end; i++ {
		registers[tokens[i]] = i
	}
}

func Lookup(ident string) Token {
	if tok, is_opcode := opcodes[ident]; is_opcode {
		return tok
	}
	if tok, is_reg := registers[ident]; is_reg {
		return tok
	}
	return ILLEGAL
}
