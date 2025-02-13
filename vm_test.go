package tmach

import (
	"math/big"
	"testing"
)

// TestLoadStore tests the LOAD and STORE instructions.
func TestLoadStore(t *testing.T) {
	vm := NewVM()

	// Set address register A0 to 0x1000
	vm.A[0] = 0x1000

	// Store a value in R0
	vm.R[0].SetInt64(123456789)
	vm.Store(0, 0) // Store R0 into memory at A0

	// Load the value back into R1
	vm.Load(1, 0) // Load from memory at A0 into R1

	// Verify the value
	if vm.R[1].Cmp(vm.R[0]) != 0 {
		t.Errorf("LOAD/STORE failed: expected %v, got %v", vm.R[0], vm.R[1])
	}
}

// TestAdd tests the ADD instruction.
func TestAdd(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(10)
	vm.R[2].SetInt64(20)

	// Perform addition
	vm.Add(0, 1, 2) // R0 = R1 + R2

	// Verify the result
	expected := big.NewInt(30)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("ADD failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestSub tests the SUB instruction.
func TestSub(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(20)
	vm.R[2].SetInt64(10)

	// Perform subtraction
	vm.Sub(0, 1, 2) // R0 = R1 - R2

	// Verify the result
	expected := big.NewInt(10)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("SUB failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestMul tests the MUL instruction.
func TestMul(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(10)
	vm.R[2].SetInt64(20)

	// Perform multiplication
	vm.Mul(0, 1, 2) // R0 = R1 * R2

	// Verify the result
	expected := big.NewInt(200)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("MUL failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestDiv tests the DIV instruction.
func TestDiv(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(20)
	vm.R[2].SetInt64(10)

	// Perform division
	vm.Div(0, 1, 2) // R0 = R1 / R2

	// Verify the result
	expected := big.NewInt(2)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("DIV failed: expected %v, got %v", expected, vm.R[0])
	}
}

func TestMod(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(20)
	vm.R[2].SetInt64(6)

	// Perform modulo operation
	vm.Mod(0, 1, 2) // R0 = R1 % R2

	// Verify the result
	expected := big.NewInt(2)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("MOD failed: expected %v, got %v", expected, vm.R[0])
	}

	// Test division by zero
	vm.R[2].SetInt64(0)
	vm.Mod(0, 1, 2) // R0 = R1 % R2 (division by zero)

	// Verify Divide-by-Zero Flag (DF) is set
	if !vm.GetFlag(DF) {
		t.Error("MOD failed: expected DF flag to be set for division by zero")
	}
}

// TestCmp tests the CMP instruction.
func TestCmp(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(10)
	vm.R[2].SetInt64(20)

	// Perform comparison
	vm.Compare(1, 2) // Compare R1 and R2

	// Verify the flags
	if !vm.GetFlag(LT) {
		t.Error("CMP failed: expected LT flag to be set")
	}
	if vm.GetFlag(GT) {
		t.Error("CMP failed: expected GT flag to be clear")
	}
	if vm.GetFlag(ZF) {
		t.Error("CMP failed: expected ZF flag to be clear")
	}
}

// TestITOF tests the ITOF instruction.
func TestITOF(t *testing.T) {
	vm := NewVM()

	// Set value in R1
	vm.R[1].SetInt64(123)

	// Perform integer to float conversion
	vm.ITOF(0, 1) // F0 = float(R1)

	// Verify the result
	expected := big.NewFloat(123)
	if vm.F[0].Cmp(expected) != 0 {
		t.Errorf("ITOF failed: expected %v, got %v", expected, vm.F[0])
	}
}

// TestFTOI tests the FTOI instruction.
func TestFTOI(t *testing.T) {
	vm := NewVM()

	// Set value in F1
	vm.F[1].SetFloat64(123.456)

	// Perform float to integer conversion
	vm.FTOI(0, 1) // R0 = int(F1)

	// Verify the result
	expected := big.NewInt(123)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("FTOI failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestAnd tests the AND instruction.
func TestAnd(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(0b1010)
	vm.R[2].SetInt64(0b1100)

	// Perform bitwise AND
	vm.And(0, 1, 2) // R0 = R1 & R2

	// Verify the result
	expected := big.NewInt(0b1000)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("AND failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestOr tests the OR instruction.
func TestOr(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(0b1010)
	vm.R[2].SetInt64(0b1100)

	// Perform bitwise OR
	vm.Or(0, 1, 2) // R0 = R1 | R2

	// Verify the result
	expected := big.NewInt(0b1110)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("OR failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestXor tests the XOR instruction.
func TestXor(t *testing.T) {
	vm := NewVM()

	// Set values in R1 and R2
	vm.R[1].SetInt64(0b1010)
	vm.R[2].SetInt64(0b1100)

	// Perform bitwise XOR
	vm.Xor(0, 1, 2) // R0 = R1 ^ R2

	// Verify the result
	expected := big.NewInt(0b0110)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("XOR failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestNot tests the NOT instruction.
func TestNot(t *testing.T) {
	vm := NewVM()

	// Set value in R1
	vm.R[1].SetInt64(0b1010)

	// Perform bitwise NOT
	vm.Not(0, 1) // R0 = ~R1

	// Verify the result
	mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	expected := new(big.Int).Xor(vm.R[1], mask)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("NOT failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestLsh tests the LSH instruction.
func TestLsh(t *testing.T) {
	vm := NewVM()

	// Set value in R1
	vm.R[1].SetInt64(0b1010)

	// Perform left shift
	vm.Lsh(0, 1, 2) // R0 = R1 << 2

	// Verify the result
	expected := big.NewInt(0b101000)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("LSH failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestRsh tests the RSH instruction.
func TestRsh(t *testing.T) {
	vm := NewVM()

	// Set value in R1
	vm.R[1].SetInt64(0b1010)

	// Perform right shift
	vm.Rsh(0, 1, 2) // R0 = R1 >> 2

	// Verify the result
	expected := big.NewInt(0b10)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("RSH failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestCsh tests the CSH instruction.
func TestCsh(t *testing.T) {
	vm := NewVM()

	// Set value in R1
	vm.R[1].SetInt64(0b1010)

	// Perform cyclic shift
	vm.Csh(0, 1, 2) // R0 = R1 <<< 2

	// Verify the result
	expected := big.NewInt(0b101000)
	if vm.R[0].Cmp(expected) != 0 {
		t.Errorf("CSH failed: expected %v, got %v", expected, vm.R[0])
	}
}

// TestJump tests the JMP instruction.
func TestJump(t *testing.T) {
	vm := NewVM()

	// Set PC to 0
	vm.PC = 0

	// Perform jump
	vm.Jump(0x1000) // Jump to address 0x1000

	// Verify PC and J
	if vm.PC != 0x1000 {
		t.Errorf("JMP failed: expected PC = 0x1000, got %v", vm.PC)
	}
	if vm.J != 1 {
		t.Errorf("JMP failed: expected J = 1, got %v", vm.J)
	}
}

// TestJumpIf tests conditional jumps.
func TestJumpIf(t *testing.T) {
	vm := NewVM()

	// Set PC to 0
	vm.PC = 0

	// Set ZF flag
	vm.SetFlag(ZF, true)

	// Perform conditional jump
	vm.JumpIf(0x1000, vm.GetFlag(ZF)) // Jump if ZF is set

	// Verify PC and J
	if vm.PC != 0x1000 {
		t.Errorf("JZ failed: expected PC = 0x1000, got %v", vm.PC)
	}
	if vm.J != 1 {
		t.Errorf("JZ failed: expected J = 1, got %v", vm.J)
	}
}
