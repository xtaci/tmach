package compiler

import "strconv"

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	IDENT
	INT
	FLOAT
	STRING
	literal_end

	COLON // :
	COMMA // ,

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
	NOP
	IN
	OUT
	LD
	ST
	XOR
	ADD
	SUB
	MUL
	DIV
	JMP
	JN
	JR
	JRN
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

	NOP: "NOP",
	IN:  "IN",
	OUT: "OUT",
	LD:  "LD",
	ST:  "ST",
	XOR: "XOR",
	ADD: "ADD",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",
	JMP: "JMP",
	JN:  "JN",
	JR:  "JR",
	JRN: "JRN",
	HLT: "HLT",
}

type Token int

func (tok Token) IsOperator() bool { return opcode_beg < tok && tok < opcode_end }
func (tok Token) IsLiteral() bool  { return literal_beg < tok && tok < literal_end }
func (tok Token) IsRegister() bool { return reg_beg < tok && tok < reg_end }

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
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
	return IDENT
}
