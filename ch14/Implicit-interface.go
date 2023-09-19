package main

import "fmt"

type MyString string

func (s MyString) Area() float64 {
	return 0
}

func (s MyString) Perimeter() float64 {
	return 0
}

func main() {

	var shape Shape
	shape = MyString("Hello")
	fmt.Println("面积:", shape.Area())
}
