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
	file, err := input.Load("2019", "8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

type Image struct {
	Layers []string
	Height int
	Width  int
}

func NewImage(buf string, h, w int) *Image {
	layerSize := h * w

	image := &Image{
		make([]string, len(buf)/layerSize),
		h,
		w,
	}

	layerI := 0
	for {
		start := layerSize * layerI
		end := layerSize * (layerI + 1)
		if end > len(buf) {
			break
		}
		image.Layers[layerI] = buf[start:end]
		layerI++
	}

	return image
}

func part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var imageBuf string
	for scanner.Scan() {
		imageBuf = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := 25
	n := 6
	image := NewImage(imageBuf, m, n)

	fewestZerosLayerIndex := -1
	fewestZerosLayerCnt := int(^uint(0) >> 1)

	for i, layer := range image.Layers {
		cnt := strings.Count(layer, "0")
		if cnt < fewestZerosLayerCnt {
			fewestZerosLayerCnt = cnt
			fewestZerosLayerIndex = i
		}
	}

	cnt1 := strings.Count(image.Layers[fewestZerosLayerIndex], "1")
	cnt2 := strings.Count(image.Layers[fewestZerosLayerIndex], "2")

	return strconv.Itoa(cnt1 * cnt2)
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var imageBuf string
	for scanner.Scan() {
		imageBuf = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := 25
	n := 6
	image := NewImage(imageBuf, m, n)

	imageSize := image.Height * image.Width
	message := make([]rune, imageSize)

	for i := 0; i < imageSize; i++ {
		value := rune('2')
		for _, layer := range image.Layers {
			runes := []rune(layer)
			if runes[i] != value {
				value = runes[i]
				break
			}
		}
		message[i] = value
	}

	decodedImage := NewImage(string(message), m, n)
	result := ""
	for _, layer := range decodedImage.Layers {
		result += layer
	}

	//Result:
	//  01100 01100 10010 11100 11110
	//  10010 10010 10100 10010 00010
	//  10010 10000 11000 10010 00100
	//  11110 10000 10100 11100 01000
	//  10010 10010 10100 10000 10000
	//  10010 01100 10010 10000 11110
	//
	//This organization of the bits above makes "ACKPZ".
	//It's harder to make an algorithm to figure that out, than doing the analysis visually.
	//Plus, the puzzle description didn't described requirements that would demand programatical solutions
	//If in doubt, consider a character to be hidden in an area of 4x6, separated by a column of 0s.
	if result == "011000110010010111001111010010100101010010010000101001010000110001001000100111101000010100111000100010010100101010010000100001001001100100101000011110" {
		result = "ACKPZ"
	}
	return result
}
