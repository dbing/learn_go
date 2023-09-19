package main

import "fmt"

func main() {
	var x int = 18
	var y *int
	y = &x
	fmt.Printf("x 变量的地址：%x\n", &x)
	fmt.Printf("y 变量的指针地址：%x\n", y)
	fmt.Printf("y 变量的值：%d\n", *y)
	fmt.Printf("y 变量的类型：%T\n", y)

}
