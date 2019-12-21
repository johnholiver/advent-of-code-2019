package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "19")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	cnt := 0
	g := grid.NewGrid(50, 50)
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			p := buildComputer(program)
			p.Input.Append(x)
			p.Input.Append(y)
			p.Process()
			o := p.Output.Read()
			g.Get(x, y).Value = o
			if o == 1 {
				cnt++
			}
		}
	}
	return fmt.Sprintln(cnt)
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	g := grid.NewGrid(50, 50)
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			p := buildComputer(program)
			p.Input.Append(x)
			p.Input.Append(y)
			p.Process()
			o := p.Output.Read()
			g.Get(x, y).Value = o
		}
	}

	return "I couldn't really make an algorithm to find the result. \n" +
		"I approximated the size of the grid and started to move a probe until I found the result.\n" +
		"The result is 6840937. The approximation code can be found at Test_part2."
}
