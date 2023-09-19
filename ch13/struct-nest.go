package ch13

import "fmt"

func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	fmt.Println(p.FirstName)
}
