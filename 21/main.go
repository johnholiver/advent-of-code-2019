package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "21")
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

	p := buildComputer(program)

	sprintScript :=
		`NOT D T
NOT T J
NOT C T
AND T J
NOT A T
OR T J
WALK
`

	for _, i := range compileScript(sprintScript) {
		p.Input.Append(i)
	}

	p.Process()
	for p.Output.CanRead() {
		p.Process()
		output := p.Output.Read()
		if output > 128 {
			return fmt.Sprint(output)
		}
		fmt.Print(string(output))
	}

	return "Didn't work .-."
}

func compileScript(s string) []int {
	r := make([]int, 0)
	for _, c := range s {
		r = append(r, int(c))
	}
	return r
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

	program += ""

	return "implement me"
}
