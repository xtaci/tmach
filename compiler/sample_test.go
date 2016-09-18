package compiler

var code1 = `
	NOP
L1:
	IN R0
	XOR R1,R1
	ST R0,R1
L2:
	IN R0
	ADD R1, 1
	ST R0, R1

	LD R0, R1
	OUT R0
	SUB R1, 1
	LD R0, R1
	OUT R0
	JMP L1
`

var code2 = `
		NOP
	L:
		IN R0
		OUT R0
		JMP L
	`

var errCode1 = `
		NOP
	L1:
		IN R0
		XOR R1 R1
		ST R0,R1
	L2:	
		IN R0
		INC R1
		ST R0 R1

	OUT:
		LD R0 R1
		OUT R0
		DEC R1	
		LD R0,R1
		OUT R0
		JMP L1
	`
