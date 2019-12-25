package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine/network"
	"io"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "23")
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
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//I had a problem in my processor which doesn't stop for input, only for output. This effectively meant that
	//my NIC if it doesn't yield an output, will be stuck in a loop of reading input, and consequently, because it
	//doesn't interrupts (no output), will try to read more and more inputs until the IO interface can't yield anything
	//any more.

	router := network.NewRouter()

	for i := 0; i < 50; i++ {
		nic := network.NewController(i, program)
		router.AddNic(nic)
	}

	var broadcastPkt *network.Packet
	for {
		router.Run()
		if router.LastBroadcast != nil {
			broadcastPkt = router.LastBroadcast
			break
		}
	}

	return fmt.Sprintf("%v", broadcastPkt.Payload.([2]int)[1])
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	router := network.NewRouter()
	router.SetDebugMode(true)

	for i := 0; i < 50; i++ {
		nic := network.NewController(i, program)
		nic.SetDebugMode(true)
		router.AddNic(nic)
	}

	var netResetPkt *network.Packet
	for {
		router.Run()
		if router.NetworkReset != nil {
			if router.NetworkReset == netResetPkt {
				break
			}
			netResetPkt = router.NetworkReset
			router.NetworkReset = nil
		}
	}

	return fmt.Sprintf("%v", netResetPkt.Payload.([2]int)[1])
}
