package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func main() {

	var s Shape
	s = Circle{Radius: 3}
	fmt.Println("area:", s.Area())
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}
