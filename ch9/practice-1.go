package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 10
	result := fibonacci(n)
	fmt.Printf("斐波那契数列的第 %d 个元素是 %d\n", n, result)
}
