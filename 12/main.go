package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/timer"
	"io"
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

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
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

	ticks, _ := findUniverseOriginFast(system)

	return fmt.Sprintf("%v", ticks)
}

func findUniverseOriginFast(system System) (int, error) {
	t := timer.New("findUniverseOriginFast")
	t.Start()

	xStates := make(map[[8]int]bool)
	yStates := make(map[[8]int]bool)
	zStates := make(map[[8]int]bool)
	initialState := system.State()
	x, y, z := splitState(initialState)
	xStates[x] = true
	yStates[y] = true
	zStates[z] = true

	xFound := false
	yFound := false
	zFound := false
	for {
		system.Tick()

		state := system.State()
		x, y, z = splitState(state)
		//if !xFound {
		_, xFound = xStates[x]
		//}
		//if !yFound {
		_, yFound = yStates[y]
		//}
		//if !zFound {
		_, zFound = zStates[z]
		//}

		if xFound && yFound && zFound {
			break
		}

		if !xFound {
			xStates[x] = true
		}
		if !yFound {
			yStates[y] = true
		}
		if !zFound {
			zStates[z] = true
		}
	}

	return LCM(len(xStates), len(yStates), len(zStates)), nil
}

func splitState(initialState [24]int) ([8]int, [8]int, [8]int) {
	var cState [3][8]int
	for i := 0; i < 3; i++ {
		cState[i] = [8]int{
			initialState[0+i],
			initialState[3+i],
			initialState[6+i],
			initialState[9+i],
			initialState[12+i],
			initialState[15+i],
			initialState[18+i],
			initialState[21+i],
		}
	}
	return cState[0], cState[1], cState[2]
}

//https://play.golang.org/p/SmzvkDjYlb
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// NEVER USE THIS! I optimized almost to the best possible scenario, and it would never be able to end
func findUniverseOriginSlow(system System) (int, error) {
	t := timer.New("findUniverseOriginSlow")
	t.Start()
	tick := 0
	initialState := system.State()
	for {
		if math.Mod(float64(tick), float64(100000000)) == 0 {
			now := time.Now()
			fmt.Printf("%v %v\n", now, tick)
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
