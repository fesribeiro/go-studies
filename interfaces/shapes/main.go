package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type retangle struct {
	width  float64
	height float64
}

func (r retangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func writeArea(s shape) {
	fmt.Printf("Area: %0.2f\n", s.area())
}

func main() {
	retangle := retangle{width: 10, height: 15}
	writeArea(retangle)
	circle := circle{radius: 10}
	writeArea(circle)
	fmt.Println("This is a simple example of interfaces in Go.")
	fmt.Println("Interfaces allow you to define methods that can be implemented by different types.")
	fmt.Println("In this example, we have a shape interface with an area method.")
}