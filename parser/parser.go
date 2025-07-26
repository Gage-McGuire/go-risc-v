package parser

import (
	lookuptables "github.com/go-risc-v/lookup-tables"
	"github.com/go-risc-v/utils"
)

func Parse(input string) utils.InstructionStructure {
	inputType := checkInputType(input)
	switch inputType {
	case "hex":
		binaryInput := hexToBinary(input)
		instructionType := extractInstructionType(binaryInput)
		instructionStructure := populateInstructionStructure(binaryInput, instructionType)
		return instructionStructure
	case "binary":
		instructionType := extractInstructionType(input)
		instructionStructure := populateInstructionStructure(input, instructionType)
		return instructionStructure
	default:
		return utils.InstructionStructure{}
	}
}

func checkInputType(input string) string {
	if len(input) > 2 && input[:2] == "0x" {
		return "hex"
	}
	for _, char := range input {
		if char != '0' && char != '1' {
			return "unknown"
		}
	}
	return "binary"
}

func hexToBinary(hex string) string {
	var binary string
	for i := 2; i < len(hex); i++ {
		if val, exists := lookuptables.HexTable[string(hex[i])]; exists {
			binary += val
		} else {
			return "invalid hex character"
		}
	}
	return binary
}

func extractInstructionType(binary string) string {
	opcode := binary[len(binary)-7:]
	if instructionType, exists := lookuptables.InstructionTypeTable[opcode]; exists {
		return instructionType
	}
	return "unknown instruction type"
}

func populateInstructionStructure(binary string, instructionType string) utils.InstructionStructure {
	switch instructionType {
	case "R-type":
		return utils.InstructionStructure{
			InstructionType: instructionType,
			Opcode:          binary[len(binary)-7:],
			Fields: utils.InstructionFields{
				Funct7: binary[0:7],
				Rs2:    binary[7:12],
				Rs1:    binary[12:17],
				Funct3: binary[17:20],
				Rd:     binary[20:25],
			},
		}
	default:
		return utils.InstructionStructure{}
	}
}
