package utils

type ConvertedInstruction struct {
	Type     string
	Mnemonic string
	Operands ConvertedOperands
}

type ConvertedOperands struct {
	Rs1 string
	Rs2 string
	Rd  string
	Imm string
}
