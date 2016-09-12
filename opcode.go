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
	CPSR // Current Processor Status Register
	REGCOUNT
)

const (
	CPSR_NEG = 1 << iota
	CPSR_ZERO
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
