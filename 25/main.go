package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := input.Load("2019", "25")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))
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

	fmt.Println("Password is 4362. Items needed: coin, hypercube, hologram, cake.")

	p.Process()

	for {
		roomOutput := ""
		for p.Output.CanRead() {
			p.Process()
			roomOutput += string(rune(p.Output.Read()))

			if strings.HasSuffix(roomOutput, "Command?") {
				break
			}
		}
		fmt.Printf("%v", roomOutput)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		for _, c := range input {
			p.Input.Append(int(c))
		}

	}
}
