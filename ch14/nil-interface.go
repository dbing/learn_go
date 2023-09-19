package main

import "fmt"

func main() {
	PrintTypeAndValue(42)
	PrintTypeAndValue("Hello")
}
func PrintTypeAndValue(val interface{}) {
	fmt.Printf("类型：%T，值：%v\n", val, val)
}
