package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/timer"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := input.Load("2019", "16")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//fmt.Printf("Result part1: %v\n", part1(file))
	//
	//file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	inputArray := toIntArray(input)

	patterns := buildPatterns(len(input))

	sum := loopChecksum(patterns, inputArray, 100)

	return printChecksum(sum)
}

const (
	limit                = 8
	realSignalMultiplier = 10000
)

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := fakeFFT(input)

	return printChecksum(sum)
}

func slowFFT(input string) []int {
	realInput := repeatSignal(input, realSignalMultiplier)
	offset := extractOffset(realInput)

	inputArray := toIntArray(realInput)

	t := timer.New("slowFFT").Start()
	patterns := buildPatterns(len(input) * realSignalMultiplier)
	fmt.Println(t.Stop().String())
	sum := loopChecksum(patterns, inputArray, 100)
	fmt.Println(t.Stop().String())
	return sum[offset : offset+limit]
}

func extractOffset(input string) int {
	offset, _ := strconv.Atoi(input[0:7])
	return offset
}

func cutInputPiece(input string, offset, limit int) string {
	return repeatSignal(input, realSignalMultiplier)[offset : offset+limit]
}

func repeatSignal(input string, multiplier int) string {
	realInput := ""
	for i := 0; i < multiplier; i++ {
		realInput += input
	}
	return realInput
}

func buildPatterns(inputLen int) [][]int {
	patterns := make([][]int, inputLen)
	basePattern := []int{0, 1, 0, -1}

	for i := 0; i < inputLen; i++ {
		patterns[i] = phasePattern(basePattern, i+1, inputLen)
	}
	return patterns
}

func buildPatternsWithLimitAndOffset(inputLen, limit, offset int) [][]int {
	patterns := make([][]int, limit)
	basePattern := []int{0, 1, 0, -1}

	for i := 0; i < limit; i++ {
		patterns[i] = phasePattern(basePattern, i+1+offset, inputLen)
	}
	return patterns
}

func phasePattern(basePattern []int, itr int, iLen int) []int {
	pLen := len(basePattern)
	pattern := make([]int, iLen+1)

	done := false
	for i := 0; i < iLen+1; i++ {
		pI := int(math.Mod(float64(i), float64(pLen)))

		for j := 0; j < itr; j++ {
			if i*itr+j > iLen {
				done = true
				break
			}
			pattern[i*itr+j] = basePattern[pI]
		}
		if done {
			break
		}
	}

	return pattern
}

func toIntArray(input string) []int {
	ss := strings.Split(input, "")
	output := make([]int, len(ss))
	for idx, s := range ss {
		output[idx], _ = strconv.Atoi(s)
	}
	return output
}

func printChecksum(input []int) string {
	s := ""
	for _, e := range input[0:8] {
		s += strconv.Itoa(e)
	}
	return s
}

func checksum(patterns [][]int, input []int) []int {
	iLen := len(input)
	output := make([]int, iLen)

	for itr := 0; itr < iLen; itr++ {
		sum := float64(0)
		for idx := 0; idx < iLen; idx++ {
			//shift 1
			shiftedIdx := idx + 1
			sum += float64(input[idx] * patterns[itr][shiftedIdx])
		}
		output[itr] = int(math.Abs(math.Mod(sum, 10)))
	}

	return output
}

func loopChecksum(patterns [][]int, input []int, n int) []int {
	for itr := 0; itr < n; itr++ {
		input = checksum(patterns, input)
	}
	return input
}

func fakeFFT(input string) []int {
	realInput := repeatSignal(input, realSignalMultiplier)
	offset := extractOffset(realInput)

	inputArray := toIntArray(realInput)

	sum := loopFakeChecksum(inputArray, 100)

	return sum[offset : offset+limit]
}

func loopFakeChecksum(input []int, n int) []int {
	for itr := 0; itr < n; itr++ {
		input = fakeChecksumSecondHalf(input)
	}
	return input
}

//Man... what the fuck... how to figure the first half? T_T
func fakeChecksumSecondHalf(input []int) []int {
	var out = make([]int, len(input))
	tmp := 0
	for i := len(input) - 1; i >= len(input)/2; i-- {
		tmp += input[i]
		out[i] = tmp % 10
	}
	return out
}
