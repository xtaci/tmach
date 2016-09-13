package compiler

import (
	"bytes"
	"encoding/binary"

	_ "github.com/xtaci/tmach"
)

func Generate(commands []interface{}) *bytes.Buffer {
	code := new(bytes.Buffer)
	for k := range commands {
		switch typedCmd := commands[k].(type) {
		case Command:
			code.WriteByte(byte(typedCmd.Op))
		case UnaryCommand:
			code.WriteByte(byte(typedCmd.Op))
			binary.Write(code, binary.LittleEndian, int32(typedCmd.X))
		case BinaryCommand:
			code.WriteByte(byte(typedCmd.Op))
			binary.Write(code, binary.LittleEndian, int32(typedCmd.X))
			binary.Write(code, binary.LittleEndian, int32(typedCmd.Y))
		}
	}
	return code
}
