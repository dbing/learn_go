package main

import (
	"fmt"
)

func findMaxAndMin(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	max := arr[0]
	min := arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return max, min
}

func main() {
	// 练习题 1: 查找数组中的最大值和最小值，
	// 编写一个函数，接受一个整数数组作为输入，并返回数组中的最大值和最小值。
	arr := []int{12, 34, 56, 7, 89, 23, 45, 68, 90, 1}
	max, min := findMaxAndMin(arr)

	fmt.Printf("最大值: %d, 最小值: %d\n", max, min)
}
