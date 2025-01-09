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

	// 6 Index register
	I0
	I1
	I2
	I3
	I4
	I5

	// JMP register
	JMPR

	// Comparision register
	CMPR
)

// Comparision register flags
const (
	EQ = 1 << iota
	GT
	LT
)

// Storage Area Length, Each unit is a number of arbitrary length
const STORAGE_SIZE = 4096
