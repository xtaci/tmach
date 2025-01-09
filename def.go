package arch

// Register IDs
const (
	// 16 Arbitrary Length Integer Register
	R0 = iota
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

	// 6 Index register(16 bit)
	I0
	I1
	I2
	I3
	I4
	I5

	// JMP register (16 bit)
	JMPR

	// Comparision register(8 bit)
	CMPR
)

// Comparision register flags
const (
	EQ = 1 << iota
	GT
	LT
)

// Area to store code
const CODE_SIZE = 65536

// Storage Area Length, Each unit is a number of arbitrary length
const STORAGE_SIZE = 4096

// Instruction Set
const (
	// Arithmetic
	ADD = iota
	SUB
	MUL
	DIV
	QUO
	MOD
	NEG

	// Logical Operation
	AND
	OR
	XOR
	NOT

	// Shift Operation
	LSH
	RSH

	// Comparision
	CMP
	CMPABS

	// Jump
	JMP
	JZ
	JNZ
	JE
	JNE
	JG
	JL
	JGE
	JLE

	// Load
	LD
	LDI // Load to Index Register

	// Store
	ST
	STI // Store Index Register

	// Extended Operation
	EXP
	MODINVERSE
	MODSQRT
	GCD
	RAND
	BINOMIAL
	SETBIT
)