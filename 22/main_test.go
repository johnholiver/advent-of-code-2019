package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_NewSmallDeck(t *testing.T) {
	deck := NewSmallDeck(6)
	assert.Equal(t, SmallDeck{0, 1, 2, 3, 4, 5}, *deck)
}

func Test_SmallDeckReverse(t *testing.T) {
	deck := NewSmallDeck(6)
	deck.DealReverse()
	assert.Equal(t, SmallDeck{5, 4, 3, 2, 1, 0}, *deck)
}

func Test_SmallDeckCutN(t *testing.T) {
	deck := NewSmallDeck(6)
	deck.CutN(3)
	assert.Equal(t, SmallDeck{3, 4, 5, 0, 1, 2}, *deck)
}

func Test_SmallDeckCutN_negative(t *testing.T) {
	deck := NewSmallDeck(6)
	deck.CutN(-2)
	assert.Equal(t, SmallDeck{4, 5, 0, 1, 2, 3}, *deck)
}

func Test_SmallDeckIncrement(t *testing.T) {
	deck := NewSmallDeck(10)
	deck.DealIncrement(3)
	assert.Equal(t, SmallDeck{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, *deck)
}

func Test_SmallSlamShuffle(t *testing.T) {
	var deck *SmallDeck
	var instructions [][2]int64
	deckSize := 10

	instructions = [][2]int64{
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal into new stack", deckSize),
		transformInstruction("deal into new stack", deckSize),
	}
	deck = NewSmallDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, SmallDeck{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}, *deck)

	instructions = [][2]int64{
		transformInstruction("cut 6", deckSize),
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal into new stack", deckSize),
	}
	deck = NewSmallDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, SmallDeck{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}, *deck)

	instructions = [][2]int64{
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal with increment 9", deckSize),
		transformInstruction("cut -2", deckSize),
	}
	deck = NewSmallDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, SmallDeck{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}, *deck)

	instructions = [][2]int64{
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
	deck = NewSmallDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, SmallDeck{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}, *deck)
}

func Test_NewBigDeck(t *testing.T) {
	deck := NewBigDeck(6)
	assert.Equal(t, big.NewInt(0), deck.offset)
	assert.Equal(t, big.NewInt(1), deck.increment)
	assert.Equal(t, big.NewInt(6), deck.size)
}

func Test_BigDeckReverse(t *testing.T) {
	deck := NewBigDeck(6)
	deck.DealReverse()
	assert.Equal(t, big.NewInt(5), deck.offset)
	assert.Equal(t, big.NewInt(5), deck.increment)
	assert.Equal(t, big.NewInt(6), deck.size)
}

func Test_BigDeckCutN(t *testing.T) {
	deck := NewBigDeck(6)
	deck.CutN(3)
	assert.Equal(t, big.NewInt(3), deck.offset)
	assert.Equal(t, big.NewInt(1), deck.increment)
	assert.Equal(t, big.NewInt(6), deck.size)
}

func Test_BigDeckCutN_negative(t *testing.T) {
	deck := NewBigDeck(6)
	deck.CutN(-2)
	assert.Equal(t, big.NewInt(4), deck.offset)
	assert.Equal(t, big.NewInt(1), deck.increment)
	assert.Equal(t, big.NewInt(6), deck.size)
}

func Test_BigDeckIncrement(t *testing.T) {
	deck := NewBigDeck(10)
	deck.DealIncrement(3)
	assert.Equal(t, big.NewInt(0), deck.offset)
	assert.Equal(t, big.NewInt(7), deck.increment)
	assert.Equal(t, big.NewInt(10), deck.size)
}

func Test_BigSlamShuffle(t *testing.T) {
	var deck *BigDeck
	var instructions [][2]int64
	deckSize := 10

	instructions = [][2]int64{
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal into new stack", deckSize),
		transformInstruction("deal into new stack", deckSize),
	}
	deck = NewBigDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, big.NewInt(0).Int64(), deck.offset.Int64())
	assert.Equal(t, big.NewInt(3), deck.increment)

	instructions = [][2]int64{
		transformInstruction("cut 6", deckSize),
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal into new stack", deckSize),
	}
	deck = NewBigDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, big.NewInt(3), deck.offset)
	assert.Equal(t, big.NewInt(7), deck.increment)

	instructions = [][2]int64{
		transformInstruction("deal with increment 7", deckSize),
		transformInstruction("deal with increment 9", deckSize),
		transformInstruction("cut -2", deckSize),
	}
	deck = NewBigDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, big.NewInt(6), deck.offset)
	assert.Equal(t, big.NewInt(7), deck.increment)

	instructions = [][2]int64{
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
	deck = NewBigDeck(deckSize)
	SlamShuffle(deck, instructions)
	assert.Equal(t, big.NewInt(9), deck.offset)
	assert.Equal(t, big.NewInt(3), deck.increment)
}
