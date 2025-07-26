package converter

import (
	lookuptables "github.com/go-risc-v/lookup-tables"
	"github.com/go-risc-v/utils"
)

func ConvertRType(instruction utils.InstructionStructure) utils.ConvertedInstruction {
	funct7 := instruction.Fields.Funct7
	funct3 := instruction.Fields.Funct3

	instructionMnemonic := lookuptables.InstructionMnemonicTable[instruction.Opcode].(map[string]any)[funct3].(map[string]string)[funct7]
	operands := convertOperands([]string{instruction.Fields.Rs1, instruction.Fields.Rs2, instruction.Fields.Rd})
	return utils.ConvertedInstruction{
		Type:     instruction.InstructionType,
		Mnemonic: instructionMnemonic,
		Operands: utils.ConvertedOperands{
			Rs1: operands[0],
			Rs2: operands[1],
			Rd:  operands[2],
		},
	}
}
