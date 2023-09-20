package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}
	rv := reflect.ValueOf(p)

	if rv.Kind() == reflect.Struct {
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			fmt.Printf("Field %d: %s = %v\n", i, rv.Type().Field(i).Name, field.Interface())
		}
	}

	// 通过反射修改值
	rv2 := reflect.ValueOf(&p).Elem()
	if rv2.Kind() == reflect.Struct {
		field := rv2.FieldByName("Age")
		if field.IsValid() && field.CanSet() {
			field.SetInt(31)
		}
	}

	fmt.Println(p) // 输出 {Alice 31}
}
