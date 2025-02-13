package tmach

// Instruction Opcodes
const (
	OP_NOP   = 0x00 // No operation
	OP_LOAD  = 0x08 // LOAD Rd, [Ax]
	OP_STORE = 0x09 // STORE Rs, [Ax]
	OP_ADD   = 0x01 // ADD Rd, Rs, Rt
	OP_SUB   = 0x02 // SUB Rd, Rs, Rt
	OP_MUL   = 0x03 // MUL Rd, Rs, Rt
	OP_DIV   = 0x04 // DIV Rd, Rs, Rt
	OP_CMP   = 0x05 // CMP Rs, Rt
	OP_ITOF  = 0x06 // ITOF Fd, Rs
	OP_FTOI  = 0x07 // FTOI Rd, Fs
	OP_AND   = 0x0A // AND Rd, Rs, Rt
	OP_OR    = 0x0B // OR Rd, Rs, Rt
	OP_XOR   = 0x0C // XOR Rd, Rs, Rt
	OP_NOT   = 0x0D // NOT Rd, Rs
	OP_LSH   = 0x0E // LSH Rd, N
	OP_RSH   = 0x0F // RSH Rd, N
	OP_CSH   = 0x10 // CSH Rd, N
	OP_JMP   = 0x11 // JMP Addr
	OP_JZ    = 0x12 // JZ Addr
	OP_JNZ   = 0x13 // JNZ Addr
	OP_JGT   = 0x14 // JGT Addr
	OP_JLT   = 0x15 // JLT Addr
	OP_JEQ   = 0x16 // JEQ Addr
)

// Status Register Flags
const (
	ZF = 0 // Zero Flag
	OF = 1 // Overflow Flag
	DF = 2 // Divide-by-Zero Flag
	LT = 3 // Less Than Flag
	GT = 4 // Greater Than Flag
)
