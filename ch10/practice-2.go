package main

import (
	"fmt"
)

func reverseArray(arr []int) []int {
	length := len(arr)
	reversed := make([]int, length)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = arr[j]
	}

	return reversed
}

func main() {
	// 练习题 2: 数组翻转
	// 编写一个函数，接受一个整数数组作为输入，并返回一个新数组，其中元素的顺序与原数组相反。
	arr := []int{1, 2, 3, 4, 5}
	reversed := reverseArray(arr)

	fmt.Println(reversed) // 输出 [5 4 3 2 1]
}
