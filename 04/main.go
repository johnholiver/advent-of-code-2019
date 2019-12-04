package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := input.Load("2019", "4")
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

	var passRange string
	for scanner.Scan() {
		passRange = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	passRangeSplit := strings.Split(passRange, "-")
	min, _ := strconv.Atoi(passRangeSplit[0])
	max, _ := strconv.Atoi(passRangeSplit[1])

	cnt := 0
	for i := min; i <= max; i++ {
		if isValid(i) {
			cnt++
		}
	}

	return strconv.Itoa(cnt)
}

func isValid(n int) bool {
	nString := strconv.Itoa(n)

	if len(nString) != 6 {
		return false
	}

	hasAdjacent := false
	for i := 0; i < len(nString)-1; i++ {
		a := nString[i : i+1]
		b := nString[i+1 : i+2]
		if a == b {
			hasAdjacent = true
		}
		if a > b {
			return false
		}
	}

	if !hasAdjacent {
		return false
	}
	return true
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)

	var passRange string
	for scanner.Scan() {
		passRange = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	passRangeSplit := strings.Split(passRange, "-")
	min, _ := strconv.Atoi(passRangeSplit[0])
	max, _ := strconv.Atoi(passRangeSplit[1])

	cnt := 0
	for i := min; i <= max; i++ {
		if isValid2(i) {
			cnt++
		}
	}

	return strconv.Itoa(cnt)
}

func isValid2(n int) bool {
	nString := strconv.Itoa(n)

	if len(nString) != 6 {
		return false
	}

	var groupMap map[string]int
	groupMap = make(map[string]int)
	for i := 0; i < len(nString)-1; i++ {
		a := nString[i : i+1]
		b := nString[i+1 : i+2]

		if a == b {
			if _, ok := groupMap[a]; !ok {
				groupMap[a] = 0
			}
			groupMap[a]++
		}

		if a > b {
			return false
		}
	}

	for _, v := range groupMap {
		if v == 1 {
			return true
		}
	}

	return false
}
