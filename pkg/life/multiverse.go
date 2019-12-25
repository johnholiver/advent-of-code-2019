package life

import "fmt"

type Multiverse struct {
	worlds   map[int]*World
	minDepth int
	maxDepth int
}

func NewMultiverse(w0 *World) *Multiverse {
	m := &Multiverse{
		make(map[int]*World),
		0,
		0,
	}
	m.addWorld(w0)
	//TODO: Not sure why I'd need to start with empty worlds on the border... but if that helps...
	m.addWorld(NewEmptyWorld(w0.width, w0.height, w0.depth-1))
	m.addWorld(NewEmptyWorld(w0.width, w0.height, w0.depth+1))
	return m
}

func (m *Multiverse) addWorld(w *World) {
	m.worlds[w.depth] = w
	w.Multiverse = m
	if w.depth < m.minDepth {
		m.minDepth = w.depth
	}
	if w.depth > m.maxDepth {
		m.maxDepth = w.depth
	}
}

func (m *Multiverse) String() string {
	mvStr := ""
	for depth := m.minDepth; depth <= m.maxDepth; depth++ {
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
		w = NewEmptyWorld(5, 5, depth)
		m.addWorld(w)
	}
	return w
}

func (m *Multiverse) Tick() {
	for depth := 0; depth <= m.maxDepth; depth++ {
		m.worlds[depth] = m.worlds[depth].Tick()
	}
	for depth := -1; depth >= m.minDepth; depth-- {
		m.worlds[depth] = m.worlds[depth].Tick()
	}
}
