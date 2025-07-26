package converter

import (
	"fmt"
	"strconv"

	"github.com/go-risc-v/assembler"
	"github.com/go-risc-v/parser"
	"github.com/go-risc-v/utils"
)

func Convert(input string) string {
	parsed := parser.Parse(input)
	instruction := getConvertedInstruction(parsed)
	assembled := assembler.AssembleInstruction(instruction)
	return assembled
}

func getConvertedInstruction(parsed utils.InstructionStructure) utils.ConvertedInstruction {
	switch parsed.InstructionType {
	case "R-type":
		return ConvertRType(parsed)
	default:
		return utils.ConvertedInstruction{}
	}
}

func convertOperands(operands []string) []string {
	var converted []string
	for _, operand := range operands {
		var intOperand int
		for _, char := range operand {
			intByte, _ := strconv.Atoi(string(char))
			intOperand = intOperand*2 + intByte
		}
		converted = append(converted, fmt.Sprintf("x%d", intOperand))
	}
	return converted
}
