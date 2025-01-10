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

	// STATUS register(8 bit)
	// 1. EQ
	// 2. GT
	// 3. LT
	// 4. Overflow
	STATUS
)

// Comparision register flags
const (
	EQ       = 1
	GT       = 2
	LT       = 4
	OVERFLOW = 8
)

const WORD_SIZE = 16 // 16 bit word size

// Word Storage Unit
const MEM_UNIT = 65536 // 64k unit of words

// Number Storage Unit
const NUM_UNIT = 4096 // 4k unit of numbers

// Ports to communicate with the outside world
const (
	PORT_CHAR = iota
	PORT_NUMERIC
)

// Instruction Set
const (
	// Arithmetic
	NOP = iota
	// ADD [dest register], [src register]
	ADD
	// SUB [dest register], [src register]
	SUB
	// MUL [dest register], [src register]
	MUL
	// DIV [dest register], [src register]
	DIV
	// MOD [dest register], [src register]
	QUO
	// MOD [dest register], [src register]
	MOD
	// NEG [REG]
	NEG
	// INC [REG]
	INC
	// DEC [REG]
	DEC

	// Logical Operation
	// AND [REG], [REG]
	AND
	// OR [REG], [REG]
	OR
	// XOR [REG], [REG]
	XOR
	// NOT [REG]
	NOT

	// Shift Operation
	// LSH [REG], IMMEDIATE
	LSH
	// RSH [REG], IMMEDIATE
	RSH

	// Comparision
	// CMP [REG], [REG]
	// CMP [REG], IMMEDIATE
	CMP

	// Jump
	// JMP [MEM_ADDR]
	JMP
	JSJ
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

	// Store words from register to memory unit
	// ST [MEM_ADDR], [REG]
	ST

	// Storge JMP register to memory unit
	// STJ [MEM_ADDR]
	STJ

	// Store number from register to number unit
	// STNUM [NUM_ADDR], [REG]
	STNUM

	// ENT
	// 1. ENT [REG], IMMEDIATE
	ENT

	// MOVE Operation
	// I0: Source
	// I1: Destination
	// I2: Length
	MOVE

	// Yield
	// Suspends the current vm and returns the control to the parent
	YIELD

	// IO Operation
	// 1. OUT REG, PORT_CHAR
	// 2. OUT REG, PORT_NUMERIC
	OUT
	// Extended Math Operation
	EXP
	MODINVERSE
	MODSQRT
	GCD
	RAND
	BINOMIAL
	SETBIT
)
