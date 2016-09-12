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
	IN  = iota // IN REGn
	OUT        // OUT REGn
	LD         // LD [REGn], REGm
	ST         // ST REGn, [REGm]
	LDR        // LDR IMM, REGn
	XOR        // XOR REGn, REGm
	ADD        // ADD REGn, REGm
	SUB        // SUB REGn, REGm
	MUL        // MUL REGn, REGm
	DIV        // DIV REGn, REGm
	INC        // INC REGn
	DEC        // DEC REGn
	JMP        // JMP IMM
	JZ         // JZ IMM
	JGZ        // JGZ IMM
	JLZ        // JLZ IMM
	HLT
)
