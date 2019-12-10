package grid

type Grid struct {
	matrix map[int]map[int]*ValuedPoint
	Width  int
	Height int
}

func NewGrid(w, h int) *Grid {
	grid := &Grid{
		Width:  w,
		Height: h,
	}
	grid.matrix = make(map[int]map[int]*ValuedPoint, h)
	for j := 0; j < h; j++ {
		grid.matrix[j] = make(map[int]*ValuedPoint, w)
		for i := 0; i < w; i++ {
			grid.matrix[j][i] = NewValuedPoint(i, j, 0)
		}
	}
	return grid
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
