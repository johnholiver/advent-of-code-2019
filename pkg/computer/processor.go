package computer

type Processor struct {
	Input         IO
	Output        IO
	Memory        Memory
	PC            int
	RelativeAddr  int
	WaitingInput  bool
	isInterrupted bool
	IsHalted      bool
}

func NewProcessor(input IO, output IO, m Memory) *Processor {
	return &Processor{
		input,
		output,
		m,
		0,
		0,
		false,
		true,
		false,
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
		if !p.Input.CanRead() {
			p.WaitingInput = true
			break
		}
		p.Memory.Write(p.PC+1, p.Input.Read(), instr.pMode[0])
	case OpcodeOutput:
		p.Output.Append(p.Memory.Read(p.PC+1, instr.pMode[0]))
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
	case OpcodeAdjustRelative:
		op1 := p.Memory.Read(p.PC+1, instr.pMode[0])
		p.RelativeAddr += op1
	case OpcodeHalt:
		p.Halt()
	}
}

func (p *Processor) Interrupt() {
	p.isInterrupted = true
}

func (p *Processor) Halt() {
	p.IsHalted = true
}

func (p *Processor) Process() error {
	p.WaitingInput = false
	p.isInterrupted = false
	for {
		instr := p.GetInstruction()
		p.ExecInstruction(instr)
		if p.WaitingInput {
			break
		}
		p.NextInstruction(instr)
		if p.isInterrupted || p.IsHalted {
			break
		}
	}
	return nil
}
