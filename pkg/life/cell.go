package life

import "fmt"

type Cell interface {
	fmt.Stringer
	HasBug() bool
	CountNeighbors() int
}

type MonoverseCell struct {
	Bug bool
	X   int
	Y   int
	W   *World
}

var _ Cell = &MonoverseCell{}

func NewMonoverseCell(bug bool, x, y int, w *World) Cell {
	return &MonoverseCell{
		bug,
		x,
		y,
		w,
	}
}

func (c *MonoverseCell) String() string {
	if c.Bug {
		return fmt.Sprintf("#")
	}
	return fmt.Sprintf(".")
}

func (c *MonoverseCell) HasBug() bool {
	return c.Bug
}

func (c *MonoverseCell) CountNeighbors() int {
	around := 0
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if nc, _ := c.W.GetCell(c.X+offset[0], c.Y+offset[1]); nc != nil && nc.HasBug() == true {
			around++
		}
	}
	return around
}

type MultiverseCell struct {
	C *MonoverseCell
}

var _ Cell = &MultiverseCell{}

func NewMultiverseCell(bug bool, x, y int, w *World) Cell {
	monoC := NewMonoverseCell(bug, x, y, w)
	return &MultiverseCell{
		monoC.(*MonoverseCell),
	}
}

func (c *MultiverseCell) String() string {
	if c.C.X == 2 && c.C.Y == 2 {
		return fmt.Sprintf("?")
	}
	return c.C.String()
}

func (c *MultiverseCell) HasBug() bool {
	return c.C.HasBug()
}

func (c *MultiverseCell) CountNeighbors() int {
	around := 0
	for _, offset := range [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		w := c.C.W
		nc, _ := w.GetCell(c.C.X+offset[0], c.C.Y+offset[1])
		if nc != nil {
			mvC := nc.(*MultiverseCell)
			if mvC.C.X == 2 && mvC.C.Y == 2 {
				//Special count of 5 cells on one of the sides of the MV depth +1
				inW := w.Multiverse.GetWorld(w.depth + 1)
				if inW == nil {
					continue
				}

				inOffsets := map[[2]int][][2]int{
					{-1, 0}: {{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}}, //right border
					{1, 0}:  {{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}, //left border
					{0, -1}: {{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}}, //down border
					{0, 1}:  {{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, //up border
				}
				for _, inOffset := range inOffsets[offset] {
					if nc, _ = inW.GetCell(inOffset[0], inOffset[1]); nc != nil && nc.HasBug() {
						around++
					}
				}
				continue
			}

			if nc.HasBug() == true {
				around++
			}
		} else {
			//Special count of 1 big cell on one of the sides of the center of the MV depth -1
			outW := w.Multiverse.GetWorld(w.depth - 1)
			if outW == nil {
				continue
			}
			if nc, _ = outW.GetCell(2+offset[0], 2+offset[1]); nc != nil && nc.HasBug() {
				around++
			}
		}
	}
	return around
}
