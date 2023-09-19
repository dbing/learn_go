package main

import "fmt"

func main() {

	var age *int

	fmt.Printf("age:%v\n", age)
	fmt.Printf("age value:%v\n", *age)
	fmt.Println(age == nil)
}
