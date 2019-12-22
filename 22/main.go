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
	file, err := input.Load("2019", "22")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	deckSize := 10007
	scanner := bufio.NewScanner(file)
	var instructions [][2]int
	for scanner.Scan() {
		instructions = append(instructions, transformInstruction(scanner.Text(), deckSize))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	deck := NewDeck(deckSize)
	SlamShuffle(deck, instructions)

	pos2019 := -1
	for i, v := range deck {
		if v == 2019 {
			pos2019 = i
			break
		}
	}

	return fmt.Sprint(pos2019)
}

func transformInstruction(instr string, deckSize int) [2]int {
	var instrInt [2]int
	instrs := strings.Split(instr, " ")
	switch instrs[0] {
	case "cut":
		arg, _ := strconv.Atoi(instrs[1])
		if arg < 0 {
			arg += deckSize
		}
		instrInt[0] = 0
		instrInt[1] = arg
	case "deal":
		switch instrs[1] {
		case "with":
			arg, _ := strconv.Atoi(instrs[3])
			instrInt[0] = 1
			instrInt[1] = arg
		case "into":
			instrInt[0] = 2
		}
	default:
		panic("No known instruction")
	}
	return instrInt
}

type Deck []int

func NewDeck(size int) Deck {
	deck := make([]int, size)
	for i, _ := range deck {
		deck[i] = i
	}
	return deck
}

func (d Deck) DealReverse() {
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
}

func (d Deck) CutN(c int) {
	if c < 0 {
		c += len(d)
	}
	aux := append(d[c:], d[:c]...)
	d.replace(aux)
}

func (d Deck) replace(aux []int) {
	for i := 0; i < len(d); i++ {
		d[i] = aux[i]
	}
}

func (d Deck) DealIncrement(inc int) {
	aux := make([]int, len(d))
	for i, j := 0, 0; i < len(d); i, j = i+1, (j+inc)%len(d) {
		aux[j] = d[i]
	}
	d.replace(aux)
}

func SlamShuffle(deck Deck, instructions [][2]int) {
	for _, instr := range instructions {
		switch instr[0] {
		case 0:
			deck.CutN(instr[1])
		case 1:
			deck.DealIncrement(instr[1])
		case 2:
			deck.DealReverse()
		default:
			panic("No known instruction")
		}
	}
}

//It appears that I can't even allocate this amount of memory (119315717514047)
//Next idea... run the instructions backwards trying to find out what will be the destiny of pos 2020
func part2(file *os.File) string {
	deckSize := 119315717514047
	scanner := bufio.NewScanner(file)
	var instructions [][2]int
	for scanner.Scan() {
		instructions = append(instructions, transformInstruction(scanner.Text(), deckSize))
	}

	//Reverse the instructions
	for i, j := 0, len(instructions)-1; i < j; i, j = i+1, j-1 {
		instructions[i], instructions[j] = instructions[j], instructions[i]
	}

	revEngPos := 2020
	for cnt := 0; cnt < 101741582076661; cnt++ {
		revEngPos = ReverseFakeSlamShuffle(cnt, revEngPos, deckSize, instructions)
	}

	return fmt.Sprint(revEngPos)
}

func ReverseFakeSlamShuffle(cnt, pos, deckSize int, instructions [][2]int) int {
	for i, instr := range instructions {
		start := pos
		switch instr[0] {
		case 0:
			pos = reverseCut(deckSize, instr[1], pos)
		case 1:
			pos = reverseIncrement(deckSize, instr[1], pos)
		case 2:
			pos = reverseSort(deckSize, pos)
		default:
			panic("No known instruction")
		}
		if start == 2020 {
			fmt.Printf("%v|%v: %v - %v\n", cnt, i, start, pos)
		}
	}
	return pos
}

func reverseSort(len, newPos int) int {
	return len - 1 - newPos
}

// How to use it!!
// (2 + x) % 8 = 3
// var res = ReverseModulus(8,2,3); // res = 1
// or...
// (cut + oldPos) % len = newPos
func reverseCut(deckLen, cut, newPos int) int {
	if newPos >= deckLen {
		panic("Remainder cannot be greater than or equal to divisor")
	}
	if cut <= newPos {
		return newPos - cut
	}
	return deckLen + newPos - cut
}

// (inc * olfPos) % len = newPos
func reverseIncrement(deckLen, inc, newPos int) int {
	if newPos >= deckLen {
		panic("Remainder cannot be greater than or equal to divisor")
	}

	div := -1
	for x := 0; x < inc; x++ {
		if (x*deckLen+newPos)%inc == 0 {
			div = x
			break
		}
	}

	return (div*deckLen + newPos) / inc
}
