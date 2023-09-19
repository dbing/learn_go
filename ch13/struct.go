package ch13

import "fmt"

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
}

func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street:  "徐家汇",
			City:    "上海",
			Country: "中国",
		},
	}

	fmt.Println(p.FirstName)
}
