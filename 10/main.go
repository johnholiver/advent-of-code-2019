package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"log"
	"math"
	"sort"
	"strings"
	"time"
)

func main() {
	file, err := input.Load("2019", "10")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	prof1 := time.Now()
	grid := buildGrid(strings.Join(lines, "\n"))
	prof2 := time.Now()
	fmt.Println("Time to build grid: ", prof2.Sub(prof1))

	start1 := time.Now()
	vp := part1(grid)
	fmt.Printf("Result part1: %v\n", vp.Value)
	stop1 := time.Now()
	fmt.Println("Time to run part1: ", stop1.Sub(start1))

	fmt.Println(vp)

	elfbet := part2(grid, vp.Point, 200)
	fmt.Printf("Result part2: %v\n", elfbet.X*100+elfbet.Y)
	stop2 := time.Now()

	fmt.Println("Time to run part2: ", stop2.Sub(stop1), "(", stop2.Sub(start1), ")")
}

func part1(grid *grid.Grid) *grid.ValuedPoint {
	//17.156563705s
	prof1 := time.Now()
	losGrid := lineOfSight(grid)
	prof2 := time.Now()
	fmt.Println("Time to build losGrid: ", prof2.Sub(prof1))

	vp := biggestLineOfSight(losGrid)

	return vp
}

type DuoPoint struct {
	Cp *grid.Point
	Pp *grid.PolarPoint
}

func part2(g *grid.Grid, mP *grid.Point, pos int) *grid.Point {
	g.Transform(-mP.X, -mP.Y)
	g.MirrorY()

	duoPoints := make([]*DuoPoint, 0)
	for j := 0; j < g.Width; j++ {
		for i := 0; i < g.Height; i++ {
			gP := g.Get(j, i)
			if !(mP.X == j && mP.Y == i) && gP.Value == 1 {
				pp := grid.NewPolarPoint(gP.Point)
				dp := &DuoPoint{grid.NewPoint(j, i), pp}
				fmt.Println(dp)
				duoPoints = append(duoPoints, dp)
			}
		}
	}

	for _, dp := range duoPoints {
		dp.Pp.Rotate(float64(-90))
	}

	sort.Slice(duoPoints, func(i, j int) bool {
		return duoPoints[i].Pp.Angle > duoPoints[j].Pp.Angle ||
			(duoPoints[i].Pp.Angle == duoPoints[j].Pp.Angle && duoPoints[i].Pp.Ro < duoPoints[j].Pp.Ro)
	})

	fmt.Println("Ordered")
	for _, dp := range duoPoints {
		fmt.Println(dp)
	}

	head := make([]*DuoPoint, 0, len(duoPoints))
	head = append(head, duoPoints[0])

	tails := make([][]*DuoPoint, 0)

	last := duoPoints[0].Pp.Angle
	lastCopyCnt := 0
	for _, dp := range duoPoints[1:] {
		if last != dp.Pp.Angle {
			head = append(head, dp)
			lastCopyCnt = 0
			last = dp.Pp.Angle
		} else {
			if lastCopyCnt == len(tails) {
				tails = append(tails, make([]*DuoPoint, 0))
			}
			tails[lastCopyCnt] = append(tails[lastCopyCnt], dp)
			lastCopyCnt++
		}
	}

	for _, tail := range tails {
		head = append(head, tail...)
	}

	duoPoints = head

	fmt.Println("Ordered (2)")
	for _, dp := range duoPoints {
		fmt.Println(dp)
	}

	return duoPoints[pos-1].Cp
}

func buildGrid(input string) *grid.Grid {
	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0])

	g := grid.NewGrid(w, h)

	for j := 0; j < g.Height; j++ {
		line := lines[j]
		for i, c := range line {
			if string(c) == "#" {
				gvp := g.Get(i, j)
				gvp.Value = 1
			}
		}
	}
	return g
}

func lineOfSight(g *grid.Grid) *grid.Grid {
	losG := grid.NewGrid(g.Width, g.Height)
	for j := 0; j < g.Width; j++ {
		for i := 0; i < g.Height; i++ {
			losG.Get(i, j).Value = pointLos2(g, i, j)
		}
	}

	return losG
}

func pointLos2(g *grid.Grid, i int, j int) int {
	gvp := g.Get(i, j)
	if gvp.Value == 0 {
		return 0
	}
	myPoint := *gvp.Point
	otherPoints := pointsFilter(g, func(gP *grid.ValuedPoint) bool {
		return gP.Value == 1 && !gP.Point.Equals(myPoint)
	})

	losSet := make([]grid.Point, 0)

	for _, oPoint := range otherPoints {
		innerPoints := pointsFilter(g, func(gP *grid.ValuedPoint) bool {
			return gP.Value == 1 &&
				!gP.Point.Equals(myPoint) &&
				!gP.Point.Equals(oPoint) &&
				pointInLine(myPoint, oPoint, *gP.Point)
		})
		if len(innerPoints) == 0 {
			losSet = append(losSet, oPoint)
		}
	}

	return len(losSet)
}

func pointsFilter(g *grid.Grid, filteringMethod func(*grid.ValuedPoint) bool) []grid.Point {
	otherPoints := make([]grid.Point, 0)
	for j := 0; j < g.Width; j++ {
		for i := 0; i < g.Height; i++ {
			gP := g.Get(i, j)
			if filteringMethod(gP) {
				otherPoints = append(otherPoints, *gP.Point)
			}
		}
	}
	return otherPoints
}

func pointInLine(a, b, c grid.Point) bool {
	crossproduct := (c.Y-a.Y)*(b.X-a.X) - (c.X-a.X)*(b.Y-a.Y)
	// compare versus epsilon for floating point values, or != 0 if using integers
	if math.Abs(float64(crossproduct)) != 0 {
		return false
	}
	dotproduct := (c.X-a.X)*(b.X-a.X) + (c.Y-a.Y)*(b.Y-a.Y)
	if dotproduct < 0 {
		return false
	}
	squaredlengthba := (b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y)
	if dotproduct > squaredlengthba {
		return false
	}
	return true
}

func biggestLineOfSight(g *grid.Grid) *grid.ValuedPoint {
	var vp *grid.ValuedPoint
	for j := 0; j < g.Width; j++ {
		for i := 0; i < g.Height; i++ {
			gvp := g.Get(i, j)
			if vp == nil || vp.Value < gvp.Value {
				vp = gvp
			}
		}
	}
	return vp
}

//
// After battling enough, I've decided to abandon this approach midway
// Everything below this comment is trash
//

func pointLos(g *grid.Grid, i int, j int) int {
	losSet := make([]grid.Point, 0)
	gvp := g.Get(i, j)
	if gvp.Value != 0 {
		a := *gvp.Point
		fmt.Printf("Checking LOS for: (%v,%v)\n", a.X, a.Y)
		borders := g.GetBorders()
		for _, bp := range borders {
			b := bp
			losPs := grid.Bresenham(*gvp.Point, bp)
			losPs = orderByDistance(losPs, a)
			if len(losPs) > 1 {
				for _, losP := range losPs[1:] {
					gLosP := g.Get(losP.X, losP.Y)
					if gLosP.Value != 0 {
						c := *gLosP.Point
						if pointInLine(a, b, c) && !containsPoint(losSet, c) {
							fmt.Printf("B(%v,%v) => C(%v,%v)\n", b.X, b.Y, c.X, c.Y)
							losSet = addToLosSet(a, c, losSet)
							break
						}
					}
				}
			}

		}
		fmt.Printf("Count: %v %v\n", len(losSet), losSet)
	}
	return len(losSet)
}

func addToLosSet(a, c grid.Point, losSet []grid.Point) []grid.Point {
	//if another point NOT in line of sight
	pointBlockingLos := false
	for _, pointsInLos := range losSet {
		if pointInLine(a, c, pointsInLos) {
			pointBlockingLos = true
			break
		}
	}

	if pointBlockingLos {
		return losSet
	}

	losSet = append(losSet, c)
	return losSet
}

func orderByDistance(ps []grid.Point, a grid.Point) []grid.Point {
	sort.Slice(ps, func(i, j int) bool { return ps[i].Distance(a) < ps[j].Distance(a) })
	return ps
}

func containsPoint(s []grid.Point, e grid.Point) bool {
	for _, a := range s {
		if a.Equals(e) {
			return true
		}
	}
	return false
}
