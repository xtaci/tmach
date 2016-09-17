package arch

// general purpose register
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
	PSR // program status register
	PC
	REGISTER_COUNT
)

const (
	// The Program Status Registers (CPSR and SPSRs)
	PSR_NEG      = 1 << 31 // Negative result from ALU flag
	PSR_ZERO     = 1 << 30 // Zero result from ALU flag.
	PSR_CARRY    = 1 << 29 // ALU operation Carried out
	PSR_OVERFLOW = 1 << 28 // ALU operation oVerflowed
)

// encoding (16b)
// |2b type| 4b operation code|1b immediate| 1b reserved | 4b Rn  |4b Rm |
const (
	TYPE_IO     = 0
	TYPE_ALU    = 1
	TYPE_MEM    = 2
	TYPE_BRANCH = 3
)

const (
	TypeShift = 14
	TypeMask  = 0x3 << TypeShift
	OpShift   = 10
	OpMask    = 0xF << OpShift
	ImmShift  = 9
	ImmMask   = 1 << ImmShift
	RnShift   = 4
	RnMask    = 0xF << RnShift
	RmMask    = 0xF
)

// IO operations
const (
	NOP = 0
	IN  = 1 // IN Rn
	OUT = 2 // OUT Rn
	HLT = 3
)

// LoadStore operations
const (
	LD = 0 // LD Rn, [Rm]/Imm
	ST = 1 // ST Rn, [Rm]/Imm
)

// Data processing
const (
	XOR = 0 // XOR Rn, Rm/Imm
	ADD = 1 // ADD Rn, Rm/Imm
	SUB = 2 // SUB Rn, Rm/Imm
	MUL = 3 // MUL Rn, Rm/Imm
	DIV = 4 // DIV Rn, Rm/Imm
)

// Branching
const (
	B   = 0 // B label
	BZ  = 1 // BZ label
	BN  = 2 // BN label
	BX  = 3 // BX Rn
	BXZ = 4 // BXZ Rn
	BXN = 5 // BXN Rn
)
