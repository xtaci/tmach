package tmach

const (
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

	R13  // R13(sp) stack pointer
	R14  // R14(lr) link register
	R15  // R15(pc) program counter
	CPSR // current program status register
	SPSR // saved program status register
	REGCOUNT
)

const (
	// The Program Status Registers (CPSR and SPSRs)
	COND_NEG      = 1 << 31 // Negative result from ALU flag
	COND_ZERO     = 1 << 30 // Zero result from ALU flag.
	COND_CARRY    = 1 << 29 // ALU operation Carried out
	COND_OVERFLOW = 1 << 28 // ALU operation oVerflowed
)

const (
	IN  = iota // IN Rn
	OUT        // OUT Rn
	LD         // LD Rn, Rm
	ST         // ST Rn, Rm
	XOR        // XOR Rn, Rm/Imm
	ADD        // ADD Rn, Rm/Imm
	SUB        // SUB Rn, Rm/Imm
	MUL        // MUL Rn, Rm/Imm
	DIV        // DIV Rn, Rm/Imm
	INC        // INC Rn
	DEC        // DEC Rn
	B          // B label
	BZ         // BZ label
	BN         // BN label
	BX         // BX Rn
	BXZ        // BXZ Rn
	BXN        // BXN Rn
	HLT
)
