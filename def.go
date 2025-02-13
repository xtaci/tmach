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
	OP_MOD   = 0x05 // MOD Rd, Rs, Rt
	OP_CMP   = 0x06 // CMP Rs, Rt
	OP_ITOF  = 0x07 // ITOF Fd, Rs
	OP_FTOI  = 0x0A // FTOI Rd, Fs
	OP_AND   = 0x0B // AND Rd, Rs, Rt
	OP_OR    = 0x0C // OR Rd, Rs, Rt
	OP_XOR   = 0x0D // XOR Rd, Rs, Rt
	OP_NOT   = 0x0E // NOT Rd, Rs
	OP_LSH   = 0x0F // LSH Rd, N
	OP_RSH   = 0x10 // RSH Rd, N
	OP_CSH   = 0x11 // CSH Rd, N
	OP_JMP   = 0x12 // JMP Addr
	OP_JZ    = 0x13 // JZ Addr
	OP_JNZ   = 0x14 // JNZ Addr
	OP_JGT   = 0x15 // JGT Addr
	OP_JLT   = 0x16 // JLT Addr
	OP_JEQ   = 0x17 // JEQ Addr
)

// Status Register Flags
const (
	ZF = 0 // Zero Flag
	OF = 1 // Overflow Flag
	DF = 2 // Divide-by-Zero Flag
	LT = 3 // Less Than Flag
	GT = 4 // Greater Than Flag
)
