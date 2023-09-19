package main

import "fmt"

type Stringer interface {
	String() string
}

type IntToStr struct {
	Value string
}

func (i IntToStr) String() string {
	return i.Value
}

func main() {
	var v Stringer
	v = IntToStr{Value: "str"}
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}
