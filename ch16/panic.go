package main

import "fmt"

func example() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了 panic:", err)
		}
	}()

	// 引发 panic
	panic("这是一个 panic")
}

func main() {
	example()
}
