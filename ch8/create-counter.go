package main

import "fmt"

func main() {
	counter := createCounter()

	// 使用闭包函数增加计数并获取计数值
	fmt.Println("计数器值:", counter()) // 输出 1
	fmt.Println("计数器值:", counter()) // 输出 2
	fmt.Println("计数器值:", counter()) // 输出 3
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
