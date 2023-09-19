package main

import (
	"fmt"
)

func removeDuplicates(slice []int) []int {
	unique := make([]int, 0, len(slice))
	seen := make(map[int]bool)

	for _, value := range slice {
		if !seen[value] {
			unique = append(unique, value)
			seen[value] = true
		}
	}

	return unique
}

func main() {
	// 习题2：编写一个函数，接受一个整数切片作为输入，并返回一个新切片，其中删除了重复的元素。
	slice := []int{1, 2, 2, 3, 4, 4, 5, 5, 5}
	uniqueSlice := removeDuplicates(slice)

	fmt.Println(uniqueSlice) // 输出 [1 2 3 4 5]
}
