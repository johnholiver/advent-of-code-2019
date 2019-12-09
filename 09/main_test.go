package main

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_io(t *testing.T) {
	program := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	i := computer_io.NewTape()
	o := computer_io.NewTape()
	p := computer.NewProcessor(i, o, nil)
	m := computer_mem.NewRelative(p, program)
	p.Memory = m
	p.Process()
	for _, op := range strings.Split(program, ",") {
		opInt, _ := strconv.Atoi(op)
		assert.Equal(t, opInt, p.Output.Read())
	}
}

func Test_16(t *testing.T) {
	program := "1102,34915192,34915192,7,4,7,99,0"
	i := computer_io.NewTape()
	o := computer_io.NewTape()
	p := computer.NewProcessor(i, o, nil)
	m := computer_mem.NewRelative(p, program)
	p.Memory = m
	p.Process()
	output := p.Output.Read()
	assert.True(t, output > 1000000000000000)
}

func Test_largeNumber(t *testing.T) {
	program := "104,1125899906842624,99"
	i := computer_io.NewTape()
	o := computer_io.NewTape()
	p := computer.NewProcessor(i, o, nil)
	m := computer_mem.NewRelative(p, program)
	p.Memory = m
	p.Process()
	assert.Equal(t, 1125899906842624, p.Output.Read())
}

func Test_part1(t *testing.T) {

}
