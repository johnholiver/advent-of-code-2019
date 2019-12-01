package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := input.Load("2019", "1")
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

	sum := 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		i = fuelPerModule(i)
		sum += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(sum)
}

func fuelPerModule(moduleWeight int) int {
	fuel := moduleWeight
	fuel /= 3
	fuel -= 2
	return fuel
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		i = recursiveFuelPerModule(i)
		sum += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(sum)
}

func recursiveFuelPerModule(moduleWeight int) int {
	totalFuel := fuelPerModule(moduleWeight)
	extraFuel := totalFuel
	for {
		extraFuel = fuelPerModule(extraFuel)
		if extraFuel <= 0 {
			break
		}
		totalFuel += extraFuel
	}
	return totalFuel
}
