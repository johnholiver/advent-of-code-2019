package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var input1 = `12345678`
var input_aoc = `59776034095811644545367793179989602140948714406234694972894485066523525742503986771912019032922788494900655855458086979764617375580802558963587025784918882219610831940992399201782385674223284411499237619800193879768668210162176394607502218602633153772062973149533650562554942574593878073238232563649673858167635378695190356159796342204759393156294658366279922734213385144895116649768185966866202413314939692174223210484933678866478944104978890019728562001417746656699281992028356004888860103805472866615243544781377748654471750560830099048747570925902575765054898899512303917159138097375338444610809891667094051108359134017128028174230720398965960712`

func Test_checksum(t *testing.T) {
	inputArray := toIntArray(input1)
	patterns := [][]int{
		{0, 1, 0, -1, 0, 1, 0, -1, 0},
		{0, 0, 1, 1, 0, 0, -1, -1, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	inputArray = checksum(patterns, inputArray)
	assert.Equal(t, "48226158", printChecksum(inputArray))

	inputArray = checksum(patterns, inputArray)
	assert.Equal(t, "34040438", printChecksum(inputArray))

	inputArray = checksum(patterns, inputArray)
	assert.Equal(t, "03415518", printChecksum(inputArray))

	inputArray = checksum(patterns, inputArray)
	assert.Equal(t, "01029498", printChecksum(inputArray))
}

func Test_fakeChecksumSecondHalf(t *testing.T) {
	inputArray := toIntArray(input1)

	inputArray = fakeChecksumSecondHalf(inputArray)
	assert.Equal(t, "00006158", printChecksum(inputArray))

	inputArray = fakeChecksumSecondHalf(inputArray)
	assert.Equal(t, "00000438", printChecksum(inputArray))

	inputArray = fakeChecksumSecondHalf(inputArray)
	assert.Equal(t, "00005518", printChecksum(inputArray))

	inputArray = fakeChecksumSecondHalf(inputArray)
	assert.Equal(t, "00009498", printChecksum(inputArray))
}

func Test_checksum100(t *testing.T) {
	inputArray := toIntArray("80871224585914546619083218645595")
	ps := buildPatterns(len(inputArray))

	sum := loopChecksum(ps, inputArray, 100)
	assert.Equal(t, "24176176", printChecksum(sum))

	inputArray = toIntArray("19617804207202209144916044189917")
	sum = loopChecksum(ps, inputArray, 100)
	assert.Equal(t, "73745418", printChecksum(sum))

	inputArray = toIntArray("69317163492948606335995924319873")
	sum = loopChecksum(ps, inputArray, 100)
	assert.Equal(t, "52432133", printChecksum(sum))
}

func Test_buildPatterns(t *testing.T) {
	ps := buildPatterns(8)
	patterns := [][]int{
		{0, 1, 0, -1, 0, 1, 0, -1, 0},
		{0, 0, 1, 1, 0, 0, -1, -1, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	assert.EqualValues(t, patterns, ps)
}

func Test_buildPatternsWithLimitAndOffset(t *testing.T) {
	offset := 2
	limit := 3
	ps := buildPatternsWithLimitAndOffset(8, limit, offset)
	patterns := [][]int{
		{0, 1, 0, -1, 0, 1, 0, -1, 0},
		{0, 0, 1, 1, 0, 0, -1, -1, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	assert.EqualValues(t, patterns[offset:offset+limit], ps)
}

func Test_cutInputPiece(t *testing.T) {
	offset := 7
	inputArray := toIntArray(repeatSignal("98765432109876543210", 1)[offset : offset+limit])
	assert.Equal(t, "21098765", printChecksum(inputArray))
}

func Test_slowFFT(t *testing.T) {
	t.Skip("This don't finish in acceptable time for full tests")

	var sum []int
	sum = slowFFT("12")
	assert.Equal(t, "00000000", printChecksum(sum))
	//1m5s
	//buildPatterns: 1.59s
	//loopChecksum: 1m3
	sum = slowFFT("00")
	assert.Equal(t, "00000000", printChecksum(sum))
	//16s
	//buildPatterns: 506.506357ms
	//loopChecksum: 14.37s
	sum = slowFFT("0")
	assert.Equal(t, "00000000", printChecksum(sum))
}

func Test_part2(t *testing.T) {
	sum := fakeFFT("03036732577212944063491565474664")
	assert.Equal(t, "84462026", printChecksum(sum))

	sum = fakeFFT("02935109699940807407585447034323")
	assert.Equal(t, "78725270", printChecksum(sum))

	sum = fakeFFT("03081770884921959731165446850517")
	assert.Equal(t, "53553731", printChecksum(sum))
}
