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

	// Program Counter
	PC
	// Stack Pointer
	SP
	// Base Pointer
	BP

	// STATUS register(8 bit)
	// 1. EQ
	// 2. GT
	// 3. LT
	// 4. Overflow
	// 5. Zero
	// 6. Negative
	STATUS
)

// Comparision register flags
const (
	EQ       = 1
	GT       = 2
	LT       = 4
	OVERFLOW = 8
	ZERO     = 16
	NEGATIVE = 32
)

// Stack Storage Unit
const STACK_UNIT = 65536

// Number Storage Unit
// 4k unit of arbitrary length numbers
const NUM_UNIT = 4096

// Assembly Opcodes
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

	// Stack Operation
	// PUSH [REG]
	PUSH
	// POP [REG]
	POP

	// Call and Return
	// CALL [MEM_ADDR]
	CALL
	// RET
	RET

	// Comparision and Testing a register
	// CMP [REG], [REG]
	// CMP [REG], IMMEDIATE
	CMP

	// Test Zero, Sign, Negative
	// TEST [REG]
	TEST

	// Jump based on the status register
	// JMP [MEM_ADDR]
	JMP
	JZ
	JNZ
	JE
	JNE
	JG
	JL
	JGE
	JLE

	// Load a number from number unit to register
	// LOAD [REG], [NUM_ADDR]
	LOAD

	// Store index register to memory unit
	// STORE [REG], [NUM_ADDR]
	STORE

	// Storge J to memory unit
	// STOREJ [MEM_ADDR]
	STOREJ

	// Enter a number directly to register
	// 1. ENT [REG], IMMEDIATE
	ENT

	// MOVE Operation
	// I0: Source
	// I1: Destination
	// I2: Length
	MOVE

	// Yield
	// Suspends current vm and returns the control to the parent
	YIELD

	// Extended Math Operation
	EXP
	MODINVERSE
	MODSQRT
	GCD
	RAND
	BINOMIAL
	SETBIT
)

// Machine Opcodes
const ()
