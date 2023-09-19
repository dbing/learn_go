package main

import "fmt"

func main() {
	var f func()
	f = myTest
	f()
}

func myTest() {
	fmt.Println("myTest func ...")
}
