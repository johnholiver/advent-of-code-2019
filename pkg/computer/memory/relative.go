package memory

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
)

type RelativeMemory struct {
	p *computer.Processor
	*SimpleMemory
}

func NewRelative(p *computer.Processor, init string) *RelativeMemory {
	return &RelativeMemory{
		p:            p,
		SimpleMemory: NewMemory(init),
	}
}

func (m *RelativeMemory) Read(address int, mode computer.ParamMode) int {
	switch mode {
	case computer.Reference:
		reference := m.Variables[address]
		return m.Variables[reference]
	case computer.Value:
		return m.Variables[address]
	case computer.Relative:
		reference := m.Variables[address]
		reference += m.p.RelativeAddr
		return m.Variables[reference]
	}
	panic("Unknown ParamMode")

}

func (m *RelativeMemory) Write(address int, value int, mode computer.ParamMode) {
	switch mode {
	case computer.Reference:
		reference := m.Variables[address]
		m.Variables[reference] = value
	case computer.Value:
		m.Variables[address] = value
	case computer.Relative:
		reference := m.Variables[address]
		reference += m.p.RelativeAddr
		m.Variables[reference] = value
	default:
		panic("Unknown ParamMode")
	}
}
