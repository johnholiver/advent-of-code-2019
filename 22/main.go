package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"math/big"
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
	var instructions [][2]int64
	for scanner.Scan() {
		instructions = append(instructions, transformInstruction(scanner.Text(), deckSize))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	deck := NewSmallDeck(deckSize)
	SlamShuffle(deck, instructions)

	pos2019 := -1
	for i, v := range *deck {
		if v == 2019 {
			pos2019 = i
			break
		}
	}

	return fmt.Sprint(pos2019)
}

func transformInstruction(instr string, deckSize int) [2]int64 {
	var instrInt [2]int64
	instrs := strings.Split(instr, " ")
	switch instrs[0] {
	case "cut":
		arg, _ := strconv.Atoi(instrs[1])
		if arg < 0 {
			arg += deckSize
		}
		instrInt[0] = 0
		instrInt[1] = int64(arg)
	case "deal":
		switch instrs[1] {
		case "with":
			arg, _ := strconv.Atoi(instrs[3])
			instrInt[0] = 1
			instrInt[1] = int64(arg)
		case "into":
			instrInt[0] = 2
		}
	default:
		panic("No known instruction")
	}
	return instrInt
}

type Deck interface {
	DealReverse()
	CutN(c int64)
	DealIncrement(inc int64)
}

type SmallDeck []int64

func NewSmallDeck(size int) *SmallDeck {
	deck := make(SmallDeck, size)
	for i, _ := range deck {
		deck[i] = int64(i)
	}
	return &deck
}

func (d SmallDeck) DealReverse() {
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
}

func (d SmallDeck) CutN(c int64) {
	if c < 0 {
		c += int64(len(d))
	}
	aux := append(d[c:], d[:c]...)
	d.replace(aux)
}

func (d SmallDeck) replace(aux []int64) {
	for i := 0; i < len(d); i++ {
		d[i] = aux[i]
	}
}

func (d SmallDeck) DealIncrement(inc int64) {
	aux := make([]int64, len(d))
	for i, j := 0, 0; i < len(d); i, j = i+1, (j+int(inc))%len(d) {
		aux[j] = d[i]
	}
	d.replace(aux)
}

func SlamShuffle(deck Deck, instructions [][2]int64) {
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
	shuffleTimes := 101741582076661
	scanner := bufio.NewScanner(file)
	var instructions [][2]int64
	for scanner.Scan() {
		instructions = append(instructions, transformInstruction(scanner.Text(), deckSize))
	}

	deck := NewBigDeck(deckSize)
	SlamShuffle(deck, instructions)

	offsetDiff := deck.offset
	incrementMul := deck.increment
	card := big.NewInt(2020)

	increment := new(big.Int).Exp(incrementMul, big.NewInt(int64(shuffleTimes)), deck.size)

	offset := new(big.Int).Mul(
		offsetDiff,
		new(big.Int).Mul(
			new(big.Int).Add(increment, new(big.Int).Sub(deck.size, big.NewInt(1))),
			new(big.Int).Exp(new(big.Int).Sub(incrementMul, big.NewInt(1)), new(big.Int).Sub(deck.size, big.NewInt(2)), deck.size),
		),
	)

	answer := new(big.Int)

	answer.Mul(increment, card)
	answer.Add(answer, offset)
	answer.Mod(answer, deck.size)

	return fmt.Sprint(answer)
}

type BigDeck struct {
	offset    *big.Int
	increment *big.Int
	size      *big.Int
}

func NewBigDeck(size int) *BigDeck {
	return &BigDeck{big.NewInt(0), big.NewInt(1), big.NewInt(int64(size))}
}

func (d *BigDeck) applyMod() {
	d.increment.Mod(d.increment, d.size)
	d.offset.Mod(d.offset, d.size)
}

func (d *BigDeck) DealReverse() {
	d.increment.Mul(d.increment, big.NewInt(-1))
	d.offset.Add(d.offset, d.increment)

	d.applyMod()
}

func (d *BigDeck) CutN(c int64) {
	if c < 0 {
		c += d.size.Int64()
	}
	d.offset.Add(d.offset, new(big.Int).Mul(d.increment, big.NewInt(c)))

	d.applyMod()
}

func (d *BigDeck) DealIncrement(inc int64) {
	a := big.NewInt(inc)

	//According to reddit (euler theorem), but didn't worked as I expected
	// For inc = 3, and deck size of 10, inv = 1, while it should be 7
	//b := new(big.Int).Sub(d.size, big.NewInt(2))
	//inv := new(big.Int).Exp(a, b, d.size)

	inv := new(big.Int).ModInverse(a, d.size)

	d.increment.Mul(d.increment, inv)

	d.applyMod()
}
