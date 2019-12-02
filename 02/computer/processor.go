package computer

type Instruction int

const (
	Sum      Instruction = 1
	Multiply Instruction = 2
	Halt     Instruction = 99
)

type Processor struct {
	Memory *Memory
	PC     int
}

func NewProcessor(m *Memory) *Processor {
	return &Processor{
		m,
		0,
	}
}

func (p *Processor) GetInstruction() Instruction {
	return Instruction(p.Memory.Variables[p.PC])
}

func (p *Processor) NextInstruction() {
	p.PC += 4
}

func (p *Processor) ExecInstruction() {
	var instr Instruction
	var addr1, addr2, addr3 int
	instr = p.GetInstruction()
	addr1 = p.Memory.Variables[p.PC+1]
	addr2 = p.Memory.Variables[p.PC+2]
	addr3 = p.Memory.Variables[p.PC+3]

	switch instr {
	case Sum:
		op1 := p.Memory.Variables[addr1]
		op2 := p.Memory.Variables[addr2]
		p.Memory.Variables[addr3] = op1 + op2
	case Multiply:
		op1 := p.Memory.Variables[addr1]
		op2 := p.Memory.Variables[addr2]
		p.Memory.Variables[addr3] = op1 * op2
	case Halt:
		//do nothing
	}
}

func (p *Processor) Process() error {
	for {
		p.ExecInstruction()
		p.NextInstruction()
		if p.GetInstruction() == Halt {
			break
		}
	}
	return nil
}
