package computer

type Processor struct {
	Input  *IO
	Output *IO
	Memory *Memory
	PC     int
}

func NewProcessor(input *IO, output *IO, m *Memory) *Processor {
	return &Processor{
		input,
		output,
		m,
		0,
	}
}

func (p *Processor) GetInstruction() Instruction {
	value := p.Memory.Read(p.PC, ParamMode(1))
	return NewInstruction(int(value))
}

func (p *Processor) NextInstruction(instr Instruction) {
	p.PC += instr.Next()
}

func (p *Processor) ExecInstruction(instr Instruction) {
	switch instr.opcode {
	case OpcodeSum:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		op2 := p.Memory.Read(p.PC+2, instr.pMode[1])
		p.Memory.Write(p.PC+3, op1+op2, instr.pMode[2])
	case OpcodeMultiply:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		op2 := p.Memory.Read(p.PC+2, instr.pMode[1])
		p.Memory.Write(p.PC+3, op1*op2, instr.pMode[2])
	case OpcodeInput:
		p.Memory.Write(p.PC+1, p.Input.Read(), instr.pMode[0])
	case OpcodeOutput:
		p.Output.Write(p.Memory.Read(p.PC+1, instr.pMode[0]))
	case OpcodeJumpTrue:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		op2 := p.Memory.Read(p.PC+2, instr.pMode[1])
		jumpTo := p.PC + instr.len()
		if op1 != 0 {
			jumpTo = op2
		}
		p.PC = jumpTo
	case OpcodeJumpFalse:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		op2 := p.Memory.Read(p.PC+2, instr.pMode[1])
		jumpTo := p.PC + instr.len()
		if op1 == 0 {
			jumpTo = op2
		}
		p.PC = jumpTo
	case OpcodeLessThen:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		op2 := p.Memory.Read(p.PC+2, instr.pMode[1])
		value := 0
		if op1 < op2 {
			value = 1
		}
		p.Memory.Write(p.PC+3, value, instr.pMode[2])
	case OpcodeEqual:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		op2 := p.Memory.Read(p.PC+2, instr.pMode[1])
		value := 0
		if op1 == op2 {
			value = 1
		}
		p.Memory.Write(p.PC+3, value, instr.pMode[2])
	case OpcodeHalt:
		//do nothing
	}
}

func (p *Processor) Process() error {
	for {
		instr := p.GetInstruction()
		p.ExecInstruction(instr)
		p.NextInstruction(instr)
		if instr.opcode == OpcodeHalt {
			break
		}
	}
	return nil
}
