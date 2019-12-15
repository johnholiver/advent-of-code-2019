package robot

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type Robot interface {
	machine.Machine
	Move(dir int) *grid.ValuedPoint
}
