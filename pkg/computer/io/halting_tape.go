package io

import "github.com/johnholiver/advent-of-code-2019/pkg/computer"

type InterruptingTape struct {
	p    *computer.Processor
	tape *Tape
}

func NewHaltingTape(p *computer.Processor) *InterruptingTape {
	return &InterruptingTape{
		p,
		NewTape(),
	}
}

var _ computer.IO = &InterruptingTape{}

func (io *InterruptingTape) Reset() {
	io.tape.Reset()
}

func (io *InterruptingTape) Previous() {
	io.tape.Previous()
}

func (io *InterruptingTape) Next() {
	io.tape.Next()
}

func (io *InterruptingTape) Read() int {
	io.p.Interrupt()
	return io.tape.Read()
}

func (io *InterruptingTape) ReadAt(index int) int {
	return io.tape.ReadAt(index)
}

func (io *InterruptingTape) Set(values []int) {
	io.tape.Set(values)
}

func (io *InterruptingTape) WriteAt(value int, index int) {
	io.tape.WriteAt(value, index)
}

func (io *InterruptingTape) Write(value int) {
	io.p.Interrupt()
	io.tape.Write(value)
}

func (io *InterruptingTape) Append(value int) {
	io.p.Interrupt()
	io.tape.Append(value)
}
