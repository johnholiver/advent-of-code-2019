package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/14/material"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "14")
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

	mt := material.NewMaterialTable()

	for scanner.Scan() {
		mTrans := material.NewMaterialTransformation(scanner.Text())
		mt[mTrans.Produces.Material] = &mTrans

		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Ore is a RAW material and we can produce infinitely without consuming anything
	mt["ORE"] = &material.MaterialTransform{
		Produces: material.MaterialCounter{"ORE", 1},
		Consumes: nil,
	}

	f := material.NewMaterialFactory(mt)

	need := material.MaterialCounter{
		Material: "FUEL",
		Count:    1,
	}
	f.ProduceRecursive(need)

	return fmt.Sprintf("%v", f.Usage("ORE"))
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)

	mt := material.NewMaterialTable()

	for scanner.Scan() {
		mTrans := material.NewMaterialTransformation(scanner.Text())
		mt[mTrans.Produces.Material] = &mTrans
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f := material.NewMaterialFactory(mt)

	collected := material.MaterialCounter{
		Material: "ORE",
		Count:    1000000000000,
	}
	f.Stock.AddMaterialCount(collected)

	fuelProduced := f.ProduceWhileStock("FUEL")

	return fmt.Sprintf("%v", fuelProduced)
}
