## **tmach Virtual Machine Instruction Set Architecture Specification**
### **1. Registers**

#### **1.1 General-Purpose Registers**
- **Width**: 256 bits.
- **Function**: Store integer or floating-point data.
- **Classification**:
  - **Integer Registers**: `R0` to `R7` for integer operations.
  - **Floating-Point Registers**: `F0` to `F7` for floating-point operations.
- **Encoding**:
  - `R0` to `R7`: `0000` to `0111`.
  - `F0` to `F7`: `1000` to `1111`.

#### **1.2 Address Registers**
- **Width**: 32 bits.
- **Function**: Used for memory address calculations.
- **Registers**: `A0` to `A7`.
- **Encoding**: Each address register is uniquely identified by a 4-bit code from `0000` to `0111`.

#### **1.3 Special-Purpose Registers**
- **Status Register (SR)**:
  - **Width**: 8 bits.
  - **Function**: Records the results of arithmetic and logical operations.
  - **Flags**:
    - **Bit 0**: Zero Flag (ZF) – Set to 1 if the result of the previous operation is zero.
    - **Bit 1**: Overflow Flag (OF) – Set to 1 if an arithmetic overflow occurs.
    - **Bit 2**: Divide-by-Zero Flag (DF) – Set to 1 if division by zero is attempted.
    - **Bits 3–5**: Comparison Result (CR) – Encodes the outcome of a `CMP` instruction (e.g., "less than," "equal," or "greater than").
    - **Bits 6–7**: Reserved for future use.

- **Program Counter (PC)**:
  - **Function**: Stores the address of the next instruction to execute.
  - **Property**: Read-only.

- **Jump Return Register (J)**:
  - **Function**: Automatically stores the return address (the address of the instruction immediately following a jump) when a jump instruction is executed.
  - **Property**: Read-only; used to facilitate function calls and returns.

---

### **2. Instruction Set**

#### **2.1 Memory Access Instructions**
- **Function**: Load data from memory into registers or store register data into memory.
- **Address Calculation**: Memory addresses are **fully determined by address registers** (no offset).
- **Instruction Format**:
  - **LOAD**: `LOAD Rd, [Ax]`
    - Loads data from the memory address stored in `Ax` into register `Rd`.
  - **STORE**: `STORE Rs, [Ax]`
    - Stores the contents of register `Rs` into the memory address stored in `Ax`.
- **Machine Code Format (16 bits)**:
  - **Opcode**: 8 bits.
  - **Target/Source Register**: 4 bits.
  - **Address Register**: 4 bits.

**Examples**:
- `LOAD R2, [A1]`:
  - Opcode: `0x08`
  - Target Register: `0010` (R2)
  - Address Register: `0001` (A1)
  - Machine Code: `0x0821`

- `STORE R3, [A2]`:
  - Opcode: `0x09`
  - Source Register: `0011` (R3)
  - Address Register: `0010` (A2)
  - Machine Code: `0x0932`

---

#### **2.2 Arithmetic Instructions**
- **Function**: Perform arithmetic operations (addition, subtraction, multiplication, division).
- **Operand Type**: Determined by the register name (`R` for integer, `F` for floating-point).
- **Instruction Format**:
  - **ADD**: `ADD Rd, Rs, Rt`
    - Adds the values in `Rs` and `Rt`, storing the result in `Rd`.
  - **SUB**: `SUB Rd, Rs, Rt`
    - Subtracts `Rt` from `Rs`, storing the result in `Rd`.
  - **MUL**: `MUL Rd, Rs, Rt`
    - Multiplies `Rs` and `Rt`, storing the result in `Rd`.
  - **DIV**: `DIV Rd, Rs, Rt`
    - Divides `Rs` by `Rt`, storing the result in `Rd`. Sets the `DF` flag if division by zero occurs.
- **Machine Code Format (24 bits)**:
  - **Opcode**: 8 bits.
  - **Source Register 1**: 4 bits.
  - **Source Register 2**: 4 bits.
  - **Destination Register**: 4 bits.
  - **Reserved**: 4 bits (unused).

**Example**:
- `ADD R0, R1, R2`:
  - Opcode: `0x01`
  - Source Register 1: `0001` (R1)
  - Source Register 2: `0010` (R2)
  - Destination Register: `0000` (R0)
  - Machine Code: `0x010012`

---

#### **2.3 Comparison Instruction**
- **Function**: Compare two register values and update the Status Register (SR).
- **Instruction Format**:
  - **CMP**: `CMP Rs, Rt`
    - Compares `Rs` and `Rt`, updating the `ZF`, `LT`, and `GT` flags.
- **Machine Code Format (16 bits)**:
  - **Opcode**: 8 bits.
  - **Source Register 1**: 4 bits.
  - **Source Register 2**: 4 bits.

**Example**:
- `CMP R0, R1`:
  - Opcode: `0x05`
  - Source Register 1: `0000` (R0)
  - Source Register 2: `0001` (R1)
  - Machine Code: `0x0501`

---

#### **2.4 Type Conversion Instructions**
- **Function**: Convert between integer and floating-point representations.
- **Instruction Format**:
  - **ITOF**: `ITOF Fd, Rs`
    - Converts the integer value in `Rs` to a floating-point value, stored in `Fd`.
  - **FTOI**: `FTOI Rd, Fs`
    - Converts the floating-point value in `Fs` to an integer value, stored in `Rd`.

---

#### **2.5 Logical and Bitwise Instructions**
- **Function**: Perform bitwise operations.
- **Instruction Format**:
  - **AND**: `AND Rd, Rs, Rt`
    - Bitwise AND of `Rs` and `Rt`, stored in `Rd`.
  - **OR**: `OR Rd, Rs, Rt`
    - Bitwise OR of `Rs` and `Rt`, stored in `Rd`.
  - **XOR**: `XOR Rd, Rs, Rt`
    - Bitwise XOR of `Rs` and `Rt`, stored in `Rd`.
  - **NOT**: `NOT Rd, Rs`
    - Bitwise NOT of `Rs`, stored in `Rd`.
  - **LSH**: `LSH Rd, N`
    - Left-shifts the value in `Rd` by `N` bits.
  - **RSH**: `RSH Rd, N`
    - Right-shifts the value in `Rd` by `N` bits.
  - **CSH**: `CSH Rd, N`
    - Cyclically shifts the value in `Rd` by `N` bits.

---

#### **2.6 Control Flow Instructions**
- **Function**: Modify program flow based on the Status Register (SR) flags.
- **Instruction Format**:
  - **JMP**: `JMP Addr`
    - Unconditionally jumps to the address `Addr` and saves the return address in `J`.
  - **Conditional Jumps**:
    - **JZ**: Jump if `ZF == 1`.
    - **JNZ**: Jump if `ZF == 0`.
    - **JGT**: Jump if `GT == 1`.
    - **JLT**: Jump if `LT == 1`.
    - **JEQ**: Jump if `ZF == 1`.

---

#### **2.7 Miscellaneous Instruction**
- **NOP**: No operation (used for timing or alignment).

---

### **3. Machine Code Format Overview**

#### **3.1 Memory Access Instructions (16 bits)**
| Field           | Length | Description                     |
|------------------|--------|---------------------------------|
| Opcode           | 8 bits | e.g., `0x08` for `LOAD`         |
| Target/Source Reg| 4 bits | General-purpose register (R/F)  |
| Address Register | 4 bits | Address register (A0–A7)        |

#### **3.2 Arithmetic/Logical Instructions (24 bits)**
| Field           | Length | Description                     |
|------------------|--------|---------------------------------|
| Opcode           | 8 bits | e.g., `0x01` for `ADD`          |
| Source Register 1| 4 bits | e.g., `Rs`                      |
| Source Register 2| 4 bits | e.g., `Rt`                      |
| Destination Reg  | 4 bits | e.g., `Rd`                      |
| Reserved         | 4 bits | Unused                          |

#### **3.3 Jump Instructions (32 bits)**
| Field           | Length | Description                     |
|------------------|--------|---------------------------------|
| Opcode           | 8 bits | e.g., `0x10` for `JMP`          |
| Address          |32 bits | Full memory address to jump to  |

---

### **4. Offset Support at the Assembly Language Level**
- **Implementation**:
  - **Address Calculation Macros**: For example, `LOAD R0, [A1 + 30]` is translated by the assembler into two instructions:
    1. Use `ADD` to compute `A1 + 30` and store the result in a temporary address register (e.g., `A2`).
    2. Execute `LOAD R0, [A2]`.
  - **Pseudo-Instructions**: Provide syntax like `LOAD R0, [A1 + 30]`, which the assembler automatically converts into low-level instructions.

---

### **5. Summary**
The **tmach Virtual Machine Instruction Set** is designed for simplicity and flexibility, with memory addressing fully controlled by 32-bit address registers. By removing offsets from hardware instructions, the design achieves greater compactness, while offset support is reintroduced at the assembly language level through macros or pseudo-instructions. This architecture is well-suited for high-precision computations, embedded systems, and scenarios requiring efficient control flow. Future enhancements could include stack support for nested function calls and expanded status flag definitions.
