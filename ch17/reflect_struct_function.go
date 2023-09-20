package main

import (
	"fmt"
	"reflect"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{3, 4}
	rv := reflect.ValueOf(r)

	if rv.Kind() == reflect.Struct {
		method := rv.MethodByName("Area")
		if method.IsValid() {
			result := method.Call(nil)
			fmt.Printf("调用方法 Area 结果: %v\n", result[0].Interface())
		}
	}
}
