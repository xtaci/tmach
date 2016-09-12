package compiler

import "strconv"

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT
	MINUS

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
	IN   // IN Rn
	OUT  // OUT Rn
	LD   // LD Rn, Rm
	ST   // ST Rn, Rm
	XOR  // XOR Rn, Rm/Imm
	ADD  // ADD Rn, Rm/Imm
	SUB  // SUB Rn, Rm/Imm
	MUL  // MUL Rn, Rm/Imm
	DIV  // DIV Rn, Rm/Imm
	IXOR // IXOR Rn, Imm
	IADD // IADD Rn, Imm
	ISUB // ISUB Rn, Imm
	IMUL // IMUL Rn, Imm
	IDIV // IDIV Rn, Imm
	INC  // INC Rn
	DEC  // DEC Rn
	B    // B label
	BZ   // BZ label
	BN   // BN label
	BX   // BX Rn
	BXZ  // BXZ Rn
	BXN  // BXN Rn
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

	IN:   "IN",
	OUT:  "OUT",
	LD:   "LD",
	ST:   "ST",
	XOR:  "XOR",
	ADD:  "AND",
	SUB:  "SUB",
	MUL:  "MUL",
	DIV:  "DIV",
	IXOR: "IXOR",
	IADD: "IADD",
	ISUB: "ISUB",
	IMUL: "IMUL",
	IDIV: "IDIV",
	INC:  "INC",
	DEC:  "DEC",
	B:    "B",
	BN:   "BN",
	BZ:   "BZ",
	BX:   "BX",
	BXN:  "BXN",
	BXZ:  "BXZ",
	HLT:  "HLT",
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
	return ILLEGAL
}
