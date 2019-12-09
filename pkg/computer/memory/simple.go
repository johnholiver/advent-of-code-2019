package memory

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	"strconv"
	"strings"
)

type SimpleMemory struct {
	Variables []int
	Len       int
}

func NewMemory(init string) *SimpleMemory {
	tokens := strings.Split(init, ",")
	//TODO: Memory is 100 times bigger than the initial code (I think I can do something better about it)
	memory := &SimpleMemory{
		make([]int, 100*len(tokens)),
		len(tokens),
	}
	for i, token := range tokens {
		memory.Variables[i], _ = strconv.Atoi(token)
	}
	return memory
}

func (m *SimpleMemory) String() string {
	s := make([]string, m.Len)
	for i, v := range m.Variables {
		if i >= m.Len {
			break
		}
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

func (m *SimpleMemory) Read(address int, mode computer.ParamMode) int {
	switch mode {
	case computer.Reference:
		reference := m.Variables[address]
		return m.Variables[reference]
	case computer.Value:
		return m.Variables[address]
	default:
		panic("Unknown ParamMode")
	}
}

func (m *SimpleMemory) Write(address int, value int, mode computer.ParamMode) {
	switch mode {
	case computer.Reference:
		reference := m.Variables[address]
		m.Variables[reference] = value
	case computer.Value:
		m.Variables[address] = value
	default:
		panic("Unknown ParamMode")
	}
}
