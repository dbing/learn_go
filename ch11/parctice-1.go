package main

import (
	"fmt"
)

func sumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func main() {
	// 习题1：编写一个函数，接受一个整数切片作为输入，并返回切片中所有元素的和。

	slice := []int{1, 2, 3, 4, 5}
	total := sumSlice(slice)

	fmt.Printf("切片的和: %d\n", total)
}
