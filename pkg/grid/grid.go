package grid

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg"
)

type Grid struct {
	matrix    map[int]map[int]*ValuedPoint
	Width     int
	Height    int
	formatter pkg.InterfaceFormatter
}

func NewGrid(w, h int) *Grid {
	g := &Grid{
		Width:  w,
		Height: h,
	}
	g.matrix = make(map[int]map[int]*ValuedPoint, g.Height)
	for j := 0; j < g.Height; j++ {
		g.matrix[j] = make(map[int]*ValuedPoint, g.Width)
		for i := 0; i < g.Width; i++ {
			g.matrix[j][i] = NewValuedPoint(i, j, 0)
		}
	}
	g.formatter = gridDefaultFormatter
	return g
}

func (g *Grid) SetFormatter(formatter pkg.InterfaceFormatter) *Grid {
	g.formatter = formatter
	return g
}

func gridDefaultFormatter(e interface{}) string {
	cast := e.(int)
	return fmt.Sprintf("%v", cast)
}

func (g *Grid) Print() string {
	gPrint := NewGrid(g.Width, g.Height)
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			vp := g.Get(i, j)
			gPrint.matrix[vp.Y][vp.X] = vp
		}
	}

	gridStr := ""
	for j := 0; j < gPrint.Height; j++ {
		line := ""
		for i := 0; i < gPrint.Width; i++ {
			line += g.formatter(gPrint.Get(i, j).Value)
		}
		gridStr += line + "\n"
	}

	return gridStr
}

func (g *Grid) Get(x, y int) *ValuedPoint {
	return g.matrix[y][x]
}

func (g *Grid) GetBorders() []Point {
	borders := make([]Point, 0)
	walkedP := NewPoint(0, 0)
	vectors := []*Vector{
		NewVector("R", g.Width-1),
		NewVector("U", g.Height-1),
		NewVector("L", g.Width-1),
		NewVector("D", g.Height-2),
	}

	borders = append(borders, Point{walkedP.X, walkedP.Y})
	for _, v := range vectors {
		walker := NewWalker(walkedP, v)
		for {
			if walker.Finished() {
				break
			}
			walkedP = walker.WalkOne()
			borders = append(borders, Point{walkedP.X, walkedP.Y})
		}
	}
	return borders
}

func (g *Grid) Transform(x, y int) {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.matrix[j][i].Point.Transform(x, y)
		}
	}
}

func (g *Grid) MirrorX() {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.matrix[j][i].Point.MirrorX()
		}
	}
}

func (g *Grid) MirrorY() {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.matrix[j][i].Point.MirrorY()
		}
	}
}
