package main

import (
	"fmt"
	"math"
)

// Square contains just one field - side length
// and it also satisfies Shape interface
type Square struct {
	side float64
}

// Circle is just another Shape
// it also satisfies Shape interface
type Circle struct {
	radius float64
}

// Shape interface defines just one method - area()
type Shape interface {
	area() float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
