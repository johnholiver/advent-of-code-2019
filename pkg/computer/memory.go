package computer

import (
	"strconv"
	"strings"
)

type Memory struct {
	Variables []int
}

func NewMemory(init string) *Memory {
	tokens := strings.Split(init, ",")
	memory := &Memory{
		make([]int, len(tokens)),
	}
	for i, token := range tokens {
		memory.Variables[i], _ = strconv.Atoi(token)
	}
	return memory
}

func (m *Memory) String() string {
	s := make([]string, len(m.Variables))
	for i, v := range m.Variables {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

func (m *Memory) Read(address int, mode ParamMode) int {
	switch mode {
	case Reference:
		reference := m.Variables[address]
		return m.Variables[reference]
	case Value:
		return m.Variables[address]
	}
	panic("Unknown ParamMode")

}

func (m *Memory) Write(address int, value int, mode ParamMode) {
	switch mode {
	case Reference:
		reference := m.Variables[address]
		m.Variables[reference] = value
	case Value:
		m.Variables[address] = value
	default:
		panic("Unknown ParamMode")
	}
}
