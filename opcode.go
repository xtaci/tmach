package tmach

const (
	REG0 = iota
	REG1
	REG2
	REG4
	REG5
	REG6
	REG7
	REG8
)

const (
	IN = iota
	OUT
	LOAD // LOAD [REGn], REGm
	STOR // STOR REGn, [REGm]
	ADD  // ADD REGn, REGm
	SUB  // SUB REGn, REGm
	MUL  // MUL REGn, REGm
	DIV  // DIV REGn, REGm
	INC  // INC REGn
	DEC  // DEC REGn
	JMP  // JMP REGn
	JZ   // JZ REGn
	JGZ  // JGZ REGn
	JLZ  // JLZ REGn
)
