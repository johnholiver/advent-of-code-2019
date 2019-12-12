package physics

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid/geo3d"
)

type Object struct {
	Pos   *geo3d.Point
	Speed *geo3d.Vector
}

func NewObject(pos *geo3d.Point) *Object {
	return &Object{
		pos,
		geo3d.NewVector(*geo3d.NewPoint(0, 0, 0), 1),
	}
}

func (o Object) String() string {
	return fmt.Sprintf("Pos: %v Spd: %v", o.Pos, o.Speed)
}

func (a Object) Equals(b Object) bool {
	return a.Pos == b.Pos && a.Speed == b.Speed
}

func (o *Object) SetSpeed(v geo3d.Vector) {
	o.Speed = &v
}

func (o Object) TotalEnergy() int {
	return o.Pot() * o.Kin()
}

func (o Object) Pot() int {
	p := o.Pos.Abs()
	return p.X + p.Y + p.Z
}

func (o Object) Kin() int {
	p := o.Speed.P.Abs()
	return p.X + p.Y + p.Z
}
