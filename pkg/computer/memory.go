package computer

import (
	"strconv"
	"strings"
)

type Memory struct {
	Variables []int
	Pc        int
}

func NewMemory(init string) *Memory {
	tokens := strings.Split(init, ",")
	memory := &Memory{
		make([]int, len(tokens)),
		0,
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
