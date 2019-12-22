package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_NewDeck(t *testing.T) {
	deck := NewDeck(6)
	assert.Equal(t, Deck{0, 1, 2, 3, 4, 5}, deck)
}

func Test_DeckReverse(t *testing.T) {
	deck := NewDeck(6)
	deck.DealReverse()
	assert.Equal(t, Deck{5, 4, 3, 2, 1, 0}, deck)
}

func Test_CutN(t *testing.T) {
	deck := NewDeck(6)
	deck.CutN(3)
	assert.Equal(t, Deck{3, 4, 5, 0, 1, 2}, deck)
}

func Test_CutN_negative(t *testing.T) {
	deck := NewDeck(6)
	deck.CutN(-2)
	assert.Equal(t, Deck{4, 5, 0, 1, 2, 3}, deck)
}

func Test_DeckIncrement(t *testing.T) {
	deck := NewDeck(10)
	deck.DealIncrement(3)
	assert.Equal(t, Deck{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, deck)
}

func Test_SlamShuffle(t *testing.T) {
	var deck Deck
	var instructions [][2]int
	deckSize := 10

	instructions = [][2]int{
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal into new stack", deckSize),
		transformInstruction("deal into new stack", deckSize),
	}
	deck = NewDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, Deck{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}, deck)

	instructions = [][2]int{
		transformInstruction("cut 6", deckSize),
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal into new stack", deckSize),
	}
	deck = NewDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, Deck{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}, deck)

	instructions = [][2]int{
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal with increment 9", deckSize),
		transformInstruction("cut -2", deckSize),
	}
	deck = NewDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, Deck{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}, deck)

	instructions = [][2]int{
		transformInstruction("deal into new stack", deckSize),
		transformInstruction("cut -2", deckSize),
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("cut 8", deckSize),
		transformInstruction("cut -4", deckSize),
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("cut 3", deckSize),
		transformInstruction("deal with increment 9", deckSize),
		transformInstruction("deal with increment 3", deckSize),
		transformInstruction("cut -1", deckSize),
	}
	deck = NewDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, Deck{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}, deck)
}

func Test_reverseCut(t *testing.T) {
	assert.Equal(t, 3, reverseCut(6, 3, 0))
	assert.Equal(t, 4, reverseCut(6, 3, 1))
	assert.Equal(t, 5, reverseCut(6, 3, 2))
	assert.Equal(t, 0, reverseCut(6, 3, 3))
	assert.Equal(t, 1, reverseCut(6, 3, 4))
	assert.Equal(t, 2, reverseCut(6, 3, 5))
}

func Test_reverseInc(t *testing.T) {
	assert.Equal(t, 0, reverseIncrement(10, 3, 0))
	assert.Equal(t, 7, reverseIncrement(10, 3, 1))
	assert.Equal(t, 4, reverseIncrement(10, 3, 2))
	assert.Equal(t, 1, reverseIncrement(10, 3, 3))
	assert.Equal(t, 8, reverseIncrement(10, 3, 4))
	assert.Equal(t, 5, reverseIncrement(10, 3, 5))
	assert.Equal(t, 2, reverseIncrement(10, 3, 6))
	assert.Equal(t, 9, reverseIncrement(10, 3, 7))
	assert.Equal(t, 6, reverseIncrement(10, 3, 8))
	assert.Equal(t, 3, reverseIncrement(10, 3, 9))
}

var input_aoc = `deal into new stack
deal with increment 68
cut 4888
deal with increment 44
cut -7998
deal into new stack
cut -5078
deal with increment 26
cut 7651
deal with increment 60
cut 8998
deal into new stack
deal with increment 64
cut -8235
deal into new stack
deal with increment 9
cut -8586
deal with increment 49
cut -7466
deal with increment 66
cut -565
deal with increment 19
cut -6306
deal with increment 67
deal into new stack
cut 886
deal with increment 63
cut 3550
deal with increment 36
cut 5593
deal with increment 18
deal into new stack
deal with increment 70
deal into new stack
cut 5168
deal with increment 39
cut 7701
deal with increment 2
deal into new stack
deal with increment 45
cut 6021
deal with increment 46
cut -6927
deal with increment 49
cut 4054
deal into new stack
deal with increment 33
deal into new stack
deal with increment 11
cut -3928
deal with increment 19
deal into new stack
deal with increment 32
cut -7786
deal with increment 27
deal into new stack
deal with increment 37
cut -744
deal with increment 25
cut -98
deal with increment 61
cut 2042
deal with increment 71
cut 5761
deal with increment 6
cut -2628
deal with increment 33
cut -9509
deal with increment 16
cut 2599
deal with increment 28
cut 2767
deal into new stack
cut 3076
deal with increment 61
deal into new stack
cut 1182
deal with increment 4
cut 2274
deal into new stack
deal with increment 31
cut -5897
deal into new stack
cut -3323
deal with increment 29
cut 879
deal with increment 12
deal into new stack
deal with increment 34
cut -5755
deal with increment 59
cut 7437
deal into new stack
cut 5095
deal into new stack
cut 453
deal with increment 24
cut -3537
deal with increment 41
deal into new stack`

func Benchmark_ReverseFakeSlamShuffle(b *testing.B) {
	deckSize := 119315717514047
	instructionsStrings := strings.Split(input_aoc, "\n")
	instructions := make([][2]int, len(instructionsStrings))

	//Reverse the instructions
	for i, j := 0, len(instructionsStrings)-1; i < j; i, j = i+1, j-1 {
		instructions[i], instructions[j] = transformInstruction(instructionsStrings[j], deckSize), transformInstruction(instructionsStrings[i], deckSize)
	}

	revEngPos := 2020
	for i := 0; i < b.N; i++ {
		revEngPos = ReverseFakeSlamShuffle(i, revEngPos, deckSize, instructions)
	}
}

//func Test_mod(t *testing.T){
//	deckSize := 119315717514047
//	instructionsStrings := strings.Split(input_aoc,"\n")
//	instructions := make([][2]int, len(instructionsStrings))
//
//	//Reverse the instructions
//	for i, j := 0, len(instructionsStrings)-1; i < j; i, j = i+1, j-1 {
//		instructions[i], instructions[j] = transformInstruction(instructionsStrings[j],deckSize), transformInstruction(instructionsStrings[i],deckSize)
//	}
//
//	m := 101741582076661 * len(instructionsStrings) % 119315717514047
//	fmt.Println(m)
//}
