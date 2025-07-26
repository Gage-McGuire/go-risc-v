package assembler

import (
	"fmt"

	"github.com/go-risc-v/utils"
)

func AssembleInstruction(instruction utils.ConvertedInstruction) string {
	switch instruction.Type {
	case "R-type":
		return fmt.Sprintf("%s %s, %s, %s", instruction.Mnemonic, instruction.Operands.Rd, instruction.Operands.Rs1, instruction.Operands.Rs2)
	default:
		return "unknown instruction type"
	}
}
