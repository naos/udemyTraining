package main

import "fmt"

// Square contains just one field - side length
// and it implements Shape interface
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

// Shape interface defines just one method - area()
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
}
