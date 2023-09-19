package main

import "fmt"

func main() {

	x := 10

	closure := func() int {
		return x
	}

	result := closure()
	fmt.Println("闭包函数返回的结果是：", result)
}
