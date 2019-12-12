package main

import (
	"fmt"
	"testing"

	"github.com/johnholiver/advent-of-code-2019/pkg/grid/geo3d"
	"github.com/johnholiver/advent-of-code-2019/pkg/physics"
	"github.com/johnholiver/advent-of-code-2019/pkg/timer"
	"github.com/stretchr/testify/assert"
)

var input_part1 = `<x=14, y=9, z=14>
<x=9, y=11, z=6>
<x=-6, y=14, z=-4>
<x=4, y=-4, z=-3>
`

func Test_buildPlanets(t *testing.T) {
	assert.Equal(t, physics.NewObject(geo3d.NewPoint(14, 9, 14)), buildCelestialCorpse("<x=14, y=9, z=14>"))
	assert.Equal(t, physics.NewObject(geo3d.NewPoint(9, 11, 6)), buildCelestialCorpse("<x=9, y=11, z=6>"))
	assert.Equal(t, physics.NewObject(geo3d.NewPoint(-6, 14, -4)), buildCelestialCorpse("<x=-6, y=14, z=-4>"))
	assert.Equal(t, physics.NewObject(geo3d.NewPoint(4, -4, -3)), buildCelestialCorpse("<x=4, y=-4, z=-3>"))
}

func TestSystem_TotalEnergy(t *testing.T) {
	o1 := buildCelestialCorpse("<x=2, y=1, z=-3>")
	o1.SetSpeed(*geo3d.NewVector(*stringToPoint("<x=-3, y=-2, z=1>"), 1))

	o2 := buildCelestialCorpse("<x=1, y=-8, z= 0>")
	o2.SetSpeed(*geo3d.NewVector(*stringToPoint("<x=-1, y=1, z=3>"), 1))

	o3 := buildCelestialCorpse("<x=3, y=-6, z=1>")
	o3.SetSpeed(*geo3d.NewVector(*stringToPoint("<x=3, y=2, z=-3>"), 1))

	o4 := buildCelestialCorpse("<x=2, y=0, z=4>")
	o4.SetSpeed(*geo3d.NewVector(*stringToPoint("<x=1, y=-1, z=-1>"), 1))

	s := System{
		Objects: []*physics.Object{o1, o2, o3, o4},
		TickCnt: 0,
	}
	assert.Equal(t, 179, s.TotalEnergy())
}

func TestSystem_tick(t *testing.T) {
	s := System{
		Objects: []*physics.Object{
			buildCelestialCorpse("<x=-1, y=0, z=2>"),
			buildCelestialCorpse("<x=2, y=-10, z=-7>"),
			buildCelestialCorpse("<x=4, y=-8, z=8>"),
			buildCelestialCorpse("<x=3, y=5, z=-1>"),
		},
		TickCnt: 0,
	}

	s.Tick()
	assert.Equal(t, "Pos: (2,-1,1) Spd: (3,-1,-1)|1", s.Objects[0].String())
	assert.Equal(t, "Pos: (3,-7,-4) Spd: (1,3,3)|1", s.Objects[1].String())
	assert.Equal(t, "Pos: (1,-7,5) Spd: (-3,1,-3)|1", s.Objects[2].String())
	assert.Equal(t, "Pos: (2,2,0) Spd: (-1,-3,1)|1", s.Objects[3].String())
}

func TestSystem_tick10(t *testing.T) {
	s := System{
		Objects: []*physics.Object{
			buildCelestialCorpse("<x=-1, y=0, z=2>"),
			buildCelestialCorpse("<x=2, y=-10, z=-7>"),
			buildCelestialCorpse("<x=4, y=-8, z=8>"),
			buildCelestialCorpse("<x=3, y=5, z=-1>"),
		},
		TickCnt: 0,
	}

	for i := 0; i < 10; i++ {
		s.Tick()
	}
	assert.Equal(t, 179, s.TotalEnergy())
}

func TestSystem_part2_small(t *testing.T) {
	tmr := timer.New("part2_small")

	s := System{
		Objects: []*physics.Object{
			buildCelestialCorpse("<x=-1, y=0, z=2>"),
			buildCelestialCorpse("<x=2, y=-10, z=-7>"),
			buildCelestialCorpse("<x=4, y=-8, z=8>"),
			buildCelestialCorpse("<x=3, y=5, z=-1>"),
		},
		TickCnt: 0,
	}

	tmr.Start()
	tick, _ := findUniverseOrigin(s)
	fmt.Println(tmr.Stop())

	assert.Equal(t, float64(2772), tick)
}

//took 15m9.985113186s
func TestSystem_part2_large(t *testing.T) {
	tmr := timer.New("1000steps")

	s := System{
		Objects: []*physics.Object{
			buildCelestialCorpse("<x=-8, y=-10, z=0>"),
			buildCelestialCorpse("<x=5, y=5, z=10>"),
			buildCelestialCorpse("<x=2, y=-7, z=3>"),
			buildCelestialCorpse("<x=9, y=-8, z=-3>"),
		},
		TickCnt: 0,
	}

	tmr.Start()
	tick, _ := findUniverseOrigin(s)
	fmt.Println(tmr.Stop())

	assert.Equal(t, float64(4686774924), tick)
}

func Benchmark_part2_small(b *testing.B) {

	s := System{
		Objects: []*physics.Object{
			buildCelestialCorpse("<x=-1, y=0, z=2>"),
			buildCelestialCorpse("<x=2, y=-10, z=-7>"),
			buildCelestialCorpse("<x=4, y=-8, z=8>"),
			buildCelestialCorpse("<x=3, y=5, z=-1>"),
		},
		TickCnt: 0,
	}

	for i := 0; i < b.N; i++ {
		findUniverseOrigin(s)
	}
}
