package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = maxNum(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
}

func maxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
