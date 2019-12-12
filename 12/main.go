package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/timer"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/johnholiver/advent-of-code-2019/pkg/grid/geo3d"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/physics"
)

func main() {
	file, err := input.Load("2019", "12")
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
	system := System{
		make([]*physics.Object, 0),
		0,
	}
	for scanner.Scan() {
		system.Objects = append(system.Objects, buildCelestialCorpse(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalTicks := 1000
	for n := 0; n < totalTicks; n++ {
		system.Tick()
	}

	return strconv.Itoa(system.TotalEnergy())
}

type System struct {
	Objects []*physics.Object
	TickCnt int
}

func buildCelestialCorpse(line string) *physics.Object {
	return physics.NewObject(stringToPoint(line))
}

func stringToPoint(line string) *geo3d.Point {
	xyz := strings.Split(line[1:len(line)-1], ",")

	x, _ := strconv.Atoi(strings.Split(xyz[0], "=")[1])
	y, _ := strconv.Atoi(strings.Split(xyz[1], "=")[1])
	z, _ := strconv.Atoi(strings.Split(xyz[2], "=")[1])

	return geo3d.NewPoint(x, y, z)
}

func (s System) Tick() {
	//apply acceleration (gravity) - update velocitys
	s.applyAcceleration()

	//apply speed - update position
	s.applyVelocity()

	s.TickCnt++
}

func (s System) applyAcceleration() {
	compare := func(a, b int) int {
		if a < b {
			return 1
		}
		if a > b {
			return -1
		}
		return 0
	}

	for i := 0; i < len(s.Objects); i++ {
		for j := i + 1; j < len(s.Objects); j++ {
			src := s.Objects[i]
			dst := s.Objects[j]
			xDiff := compare(src.Pos.X, dst.Pos.X)
			yDiff := compare(src.Pos.Y, dst.Pos.Y)
			zDiff := compare(src.Pos.Z, dst.Pos.Z)
			src.Speed.P.X += xDiff
			src.Speed.P.Y += yDiff
			src.Speed.P.Z += zDiff
			dst.Speed.P.X -= xDiff
			dst.Speed.P.Y -= yDiff
			dst.Speed.P.Z -= zDiff
		}
	}
}

func (s System) applyVelocity() {
	for _, obj := range s.Objects {
		obj.Pos.X += obj.Speed.P.X
		obj.Pos.Y += obj.Speed.P.Y
		obj.Pos.Z += obj.Speed.P.Z
	}
}

func (s System) TotalEnergy() int {
	totalEnergy := 0
	for _, obj := range s.Objects {
		totalEnergy += obj.TotalEnergy()
	}
	return totalEnergy
}

func (s System) Hash() [32]byte {
	buf, _ := json.Marshal(s.Objects)
	return sha256.Sum256(buf)
}

func (s System) State() [24]int {
	var hash [24]int
	for i, obj := range s.Objects {
		hash[i*6+0] = obj.Pos.X
		hash[i*6+1] = obj.Pos.Y
		hash[i*6+2] = obj.Pos.Z
		hash[i*6+3] = obj.Speed.P.X
		hash[i*6+4] = obj.Speed.P.Y
		hash[i*6+5] = obj.Speed.P.Z
	}
	return hash
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	system := System{
		make([]*physics.Object, 0),
		0,
	}
	for scanner.Scan() {
		system.Objects = append(system.Objects, buildCelestialCorpse(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ticks, _ := findUniverseOrigin(system)

	return fmt.Sprintf("%.0f", ticks)
}

func findUniverseOrigin(system System) (float64, error) {
	t := timer.New("findUniverseOrigin")
	t.Start()
	tick := float64(0)
	initialState := system.State()
	for {
		if math.Mod(float64(tick), float64(100000000)) == 0 {
			now := time.Now()
			fmt.Printf("%v %.0f\n", now, tick)
			//if (t.Elapsed(now).Minutes()>3) {
			//	return tick, fmt.Errorf("Is taking more than 3 minutes")
			//}
		}

		system.Tick()
		tick++

		state := system.State()
		if state == initialState {
			break
		}
	}
	return tick, nil
}
