package tmach

const (
	R0 = iota
	R1
	R2
	R4
	R5
	R6
	R7
	R8
	REGNUM
)

const (
	IN  = iota // IN REGn
	OUT        // OUT REGn
	LD         // LD [REGn], REGm
	ST         // ST REGn, [REGm]
	LDR        // LDR Imm, REGn
	XOR        // XOR REGn, REGm
	ADD        // ADD REGn, REGm
	SUB        // SUB REGn, REGm
	MUL        // MUL REGn, REGm
	DIV        // DIV REGn, REGm
	INC        // INC REGn
	DEC        // DEC REGn
	JMP        // JMP REGn
	JZ         // JZ REGn
	JGZ        // JGZ REGn
	JLZ        // JLZ REGn
	HLT
)
