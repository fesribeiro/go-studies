package shapes

import (
	"math"
)

type shape interface {
	area() float64
}

type Retangle struct {
	width  float64
	height float64
}

func (r Retangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

