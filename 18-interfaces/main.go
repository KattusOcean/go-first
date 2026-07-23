package main

import (
	"fmt"
	"math"
)

// Interface
type Shape interface {
	area() float64
}

// Structures (constructors)
type Circle struct {
	radius float64
}

type Rectangle struct {
	width, heigth float64
}

// Methods (instead os using `implements` keyword on Java,
// with GO you "implement" using `(<Object>) <interface method>`)
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.heigth
}

func getArea(s Shape) float64 {
	return s.area()
}

// Main method
func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, heigth: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
