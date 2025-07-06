package shapes

import (
	"math"
	"testing"
)


func TestArea(t *testing.T) {
	t.Run("Retangle Area", func(t *testing.T) {
		r := Retangle{width: 10, height: 15}
		expected := 150.0
		if r.area() != expected {
			t.Fatalf("Expected area %f, got %f", expected, r.area())
		}
	})
	t.Run("Circle Area", func(t *testing.T) {
		c := Circle{10}
		expected := math.Pi * math.Pow(10, 2)
		if c.area() != expected {
			t.Fatalf("Expected area %f, got %f", expected, c.area())
		}
	})

}
