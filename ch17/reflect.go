package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t, v, v.Kind())
}
