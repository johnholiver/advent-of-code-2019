package computer

import "strconv"

type Opcode int

const (
	OpcodeSum            Opcode = 1
	OpcodeMultiply       Opcode = 2
	OpcodeInput          Opcode = 3
	OpcodeOutput         Opcode = 4
	OpcodeJumpTrue       Opcode = 5
	OpcodeJumpFalse      Opcode = 6
	OpcodeLessThen       Opcode = 7
	OpcodeEqual          Opcode = 8
	OpcodeAdjustRelative Opcode = 9
	OpcodeHalt           Opcode = 99
)

type ParamMode int

const (
	Reference ParamMode = iota
	Value
	Relative
)

type Instruction struct {
	opcode Opcode
	pMode  []ParamMode
}

func NewInstruction(i int) Instruction {
	iS := strconv.Itoa(i)

	reverseIndex := len(iS)
	var oS string
	if reverseIndex == 1 {
		oS = iS[reverseIndex-1 : reverseIndex]
	} else {
		oS = iS[reverseIndex-2 : reverseIndex]
	}
	reverseIndex -= 2

	o, _ := strconv.Atoi(oS)
	instruction := Instruction{
		opcode: Opcode(o),
	}

	instruction.pMode = make([]ParamMode, instruction.len()-1)

	ipmI := 0
	for ; reverseIndex > 0; reverseIndex-- {
		pMode, _ := strconv.Atoi(iS[reverseIndex-1 : reverseIndex])
		instruction.pMode[ipmI] = ParamMode(pMode)
		ipmI++
	}

	return instruction
}

func (i Instruction) len() int {
	switch i.opcode {
	case OpcodeSum, OpcodeMultiply, OpcodeLessThen, OpcodeEqual:
		return 4
	case OpcodeInput, OpcodeOutput, OpcodeAdjustRelative:
		return 2
	case OpcodeJumpTrue, OpcodeJumpFalse:
		return 3
	case OpcodeHalt:
		return 1
	}
	panic("Unknown instruction")
}

func (i Instruction) Next() int {
	switch i.opcode {
	case OpcodeJumpTrue, OpcodeJumpFalse:
		return 0
	}
	return i.len()
}
