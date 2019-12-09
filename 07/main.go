package main

import (
	"bufio"
	"fmt"
	prmt "github.com/gitchander/permutation"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := input.Load("2019", "7")
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

	tests := permute([]int{0, 1, 2, 3, 4})

	var bestSetting *ThrusterTest
	for _, t := range tests {
		t.OutputSignal = setThrusters(program, t.PhaseSettings)
		if bestSetting == nil || t.OutputSignal > bestSetting.OutputSignal {
			bestSetting = t
		}
	}

	return strconv.Itoa(bestSetting.OutputSignal)
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

	tests := permute([]int{5, 6, 7, 8, 9})

	var bestSetting *ThrusterTest
	for _, t := range tests {
		t.OutputSignal = setThrustersInFeedbackLoop(program, t.PhaseSettings)
		if bestSetting == nil || t.OutputSignal > bestSetting.OutputSignal {
			bestSetting = t
		}
	}

	return strconv.Itoa(bestSetting.OutputSignal)
}

func setThrusters(program string, phaseSettings []int) int {
	input := 0
	output := 0
	for _, phaseSetting := range phaseSettings {
		i := computer_io.NewTape()
		i.Set([]int{phaseSetting, input})
		o := computer_io.NewTape()

		m := computer_mem.NewMemory(program)
		p := computer.NewProcessor(i, o, m)
		p.Process()

		output = p.Output.Read()
		input = output
	}
	return output
}

func setThrustersInFeedbackLoop(program string, phaseSettings []int) int {
	ps := make([]*computer.Processor, len(phaseSettings))
	for index, phaseSetting := range phaseSettings {
		i := computer_io.NewTape()
		i.Set([]int{phaseSetting})
		m := computer_mem.NewMemory(program)
		ps[index] = computer.NewProcessor(i, nil, m)
		o := computer_io.NewHaltingTape(ps[index])
		ps[index].Output = o
	}

	signal := 0
	for {
		for index, _ := range phaseSettings {
			ps[index].Input.Append(signal)

			ps[index].Process()

			if ps[index].IsHalted {
				ps[index].Output.Previous()
			}
			signal = ps[index].Output.Read()
		}
		//Exit when the last processor is terminated
		if ps[len(phaseSettings)-1].IsHalted {
			break
		}
	}
	return signal
}

type ThrusterTest struct {
	PhaseSettings []int
	OutputSignal  int
}

func permute(array []int) []*ThrusterTest {
	tests := make([]*ThrusterTest, 0)

	p := prmt.New(prmt.IntSlice(array))
	for p.Next() {
		tt := &ThrusterTest{
			PhaseSettings: make([]int, len(array)),
			OutputSignal:  0,
		}
		copy(tt.PhaseSettings, array)
		tests = append(tests, tt)

	}

	return tests
}
