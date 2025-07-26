package utils

type InstructionStructure struct {
	InstructionType string
	Opcode          string
	Fields          InstructionFields
}

type InstructionFields struct {
	Funct7 string
	Rs2    string
	Rs1    string
	Funct3 string
	Rd     string
	Imm    string
}
