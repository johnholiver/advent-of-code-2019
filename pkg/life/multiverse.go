package life

import "fmt"

type Multiverse struct {
	worlds      map[int]*World
	OutterDepth int
	InnerDepth  int
}

func NewMultiverse(w0 *World) *Multiverse {
	m := &Multiverse{
		make(map[int]*World),
		0,
		0,
	}
	m.addWorld(w0)
	//Ideally, I should be able to expand the multiverse with the ticks, but I'm lazy to properly identify a border
	//expansion. 100 worlds should be enough to solve the puzzle.
	for i := 1; i < 101; i++ {
		wOutter := NewWorld(w0.width, w0.height, w0.depth-i)
		wOutter.FillEmpty(NewMultiverseCell)
		m.addWorld(wOutter)
		wInner := NewWorld(w0.width, w0.height, w0.depth+i)
		wInner.FillEmpty(NewMultiverseCell)
		m.addWorld(wInner)
	}
	return m
}

func (m *Multiverse) addWorld(w *World) {
	m.worlds[w.depth] = w
	w.Multiverse = m
	if w.depth < m.OutterDepth {
		m.OutterDepth = w.depth
	}
	if w.depth > m.InnerDepth {
		m.InnerDepth = w.depth
	}
}

func (m *Multiverse) String() string {
	mvStr := ""
	for depth := m.OutterDepth; depth <= m.InnerDepth; depth++ {
		mvStr += fmt.Sprintln("Depth", depth)
		mvStr += m.worlds[depth].String()
	}
	return mvStr
}

func (m *Multiverse) CountBugs() int {
	cnt := 0
	for _, w := range m.worlds {
		cnt += w.CountBugs()
	}
	return cnt
}

func (m *Multiverse) GetWorld(depth int) *World {
	w, found := m.worlds[depth]
	if !found {
		//See comment on constructor
		return nil
		//w = NewWorld(5, 5, depth)
		//w.FillEmpty(NewMultiverseCell)
		//m.addWorld(w)
	}
	return w
}

func (m *Multiverse) Tick() {
	worlds2 := make(map[int]*World)
	for depth := m.OutterDepth; depth <= m.InnerDepth; depth++ {
		w2 := m.worlds[depth].Tick()
		w2.Multiverse = m
		worlds2[depth] = w2
	}
	m.worlds = worlds2
}
