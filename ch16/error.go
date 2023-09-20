package main

import (
	"errors"
	"fmt"
	"math"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("负数无法求平方根")
	}
	return math.Sqrt(x), nil
}

func main() {
	num1, num2 := 10.0, 2.0
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println("除法出错:", err)
	} else {
		fmt.Println("除法结果:", result)
	}

	num3 := -9.0
	sqrtResult, sqrtErr := squareRoot(num3)
	if sqrtErr != nil {
		fmt.Println("求平方根出错:", sqrtErr)
	} else {
		fmt.Println("平方根结果:", sqrtResult)
	}
}
