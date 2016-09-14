package arch

const (
	R0 byte = iota
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

	R13  // R13(sp) stack pointer
	R14  // R14(lr) link register
	R15  // R15(pc) program counter
	CPSR // current program status register
	SPSR // saved program status register
	REGCOUNT

	SP = R13
	LR = R14
	PC = R15
)

const (
	// The Program Status Registers (CPSR and SPSRs)
	COND_NEG      = 1 << 31 // Negative result from ALU flag
	COND_ZERO     = 1 << 30 // Zero result from ALU flag.
	COND_CARRY    = 1 << 29 // ALU operation Carried out
	COND_OVERFLOW = 1 << 28 // ALU operation oVerflowed
)

// machine code
const (
	NOP  = 0
	IN   = 1  // IN Rn
	OUT  = 2  // OUT Rn
	LD   = 3  // LD Rn, [Rm]
	ST   = 4  // ST Rn, [Rm]
	XOR  = 5  // XOR Rn, Rm
	ADD  = 6  // ADD Rn, Rm
	SUB  = 7  // SUB Rn, Rm
	MUL  = 8  // MUL Rn, Rm
	DIV  = 9  // DIV Rn, Rm
	IXOR = 10 // IXOR Rn, Imm
	IADD = 11 // IADD Rn, Imm
	ISUB = 12 // ISUB Rn, Imm
	IMUL = 13 // IMUL Rn, Imm
	IDIV = 14 // IDIV Rn, Imm
	INC  = 15 // INC Rn
	DEC  = 16 // DEC Rn
	B    = 17 // B label
	BZ   = 18 // BZ label
	BN   = 19 // BN label
	BX   = 20 // BX Rn
	BXZ  = 21 // BXZ Rn
	BXN  = 22 // BXN Rn
	HLT  = 22
)
