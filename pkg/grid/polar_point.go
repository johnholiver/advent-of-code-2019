package grid

import (
	"fmt"
	"math"
)

type PolarPoint struct {
	Ro    float64
	Teta  float64
	Angle float64
}

func NewPolarPoint(p *Point) *PolarPoint {
	teta := float64(0)
	angle := float64(0)
	if p.X < 0 {
		angle += 180
	}

	if p.X != 0 {
		teta = math.Atan(float64(p.Y) / float64(p.X))
	} else {
		if p.Y < 0 {
			teta = -math.Pi / 2
		} else {
			teta = math.Pi / 2
		}
	}
	angle += teta * 180 / math.Pi

	if angle <= 0 {
		angle += 360
	}

	return &PolarPoint{
		Ro:    math.Sqrt(math.Pow(float64(p.X), float64(2)) + math.Pow(float64(p.Y), float64(2))),
		Teta:  teta,
		Angle: angle,
	}
}

func (p *PolarPoint) Rotate(angle float64) {
	p.Angle += angle
	p.Teta += math.Pi * angle / 180

	if p.Angle <= 0 {
		p.Angle += 360
	}
}

func (p PolarPoint) String() string {
	return fmt.Sprintf("(%v,%v)", p.Ro, p.Angle)
}
