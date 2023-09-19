package main

import "fmt"

func main() {
	func() {
		fmt.Println("匿名函数")
	}()

	func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2)

	sum := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(sum)
}
