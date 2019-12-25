package life

import (
	"errors"
	"math"
)

type World struct {
	cells      []Cell
	width      int
	height     int
	depth      int
	Multiverse *Multiverse
}

func NewWorld(width, height, depth int) *World {
	return &World{
		cells:  make([]Cell, width*height),
		depth:  depth,
		width:  width,
		height: height,
	}
}

func (w *World) FillEmpty(cellBuilder func(bool, int, int, *World) Cell) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			c := cellBuilder(false, x, y, w)
			w.SetCell(x, y, c)
		}
	}
}

func (w *World) String() string {
	wStr := ""
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			c, _ := w.GetCell(x, y)
			wStr += c.String()
		}
		wStr += "\n"
	}
	return wStr
}

func (w *World) SetCell(x, y int, c Cell) error {
	if x < 0 || x >= w.width {
		return errors.New("bad X")
	}
	if y < 0 || y >= w.height {
		return errors.New("bad Y")
	}

	pos := y*w.width + x
	w.cells[pos] = c
	return nil
}

func (w *World) GetCell(x, y int) (Cell, error) {
	if x < 0 || x >= w.width {
		return nil, errors.New("bad X")
	}
	if y < 0 || y >= w.height {
		return nil, errors.New("bad Y")
	}

	pos := y*w.width + x
	return w.cells[pos], nil
}

func (w *World) Tick() *World {
	w2 := NewWorld(w.width, w.height, w.depth)

	for y := 0; y < w2.height; y++ {
		for x := 0; x < w2.width; x++ {
			c, _ := w.GetCell(x, y)
			around := c.CountNeighbors()
			newHasBug := c.HasBug()
			switch c.HasBug() {
			case true:
				if around != 1 {
					newHasBug = false
				}
			case false:
				if around == 1 || around == 2 {
					newHasBug = true
				}
			}

			var nc Cell
			if _, ok := c.(*MonoverseCell); ok {
				nc = NewMonoverseCell(newHasBug, x, y, w2)
			}
			if _, ok := c.(*MultiverseCell); ok {
				nc = NewMultiverseCell(newHasBug, x, y, w2)
			}
			w2.SetCell(x, y, nc)
		}
	}
	return w2
}

func (w *World) CountBugs() int {
	cnt := 0
	for pos := 0; pos < len(w.cells); pos++ {
		if pos != 12 && w.cells[pos].HasBug() == true {
			cnt += 1
		}
	}
	return cnt
}

func (w *World) BiodiversityRating() int {
	rating := 0
	for pos := 0; pos < len(w.cells); pos++ {
		if w.cells[pos].HasBug() == true {
			rating += int(math.Pow(2, float64(pos)))
		}
	}
	return rating
}
