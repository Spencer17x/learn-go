package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with agrs "+"(%d, %d)", opName, a, b)
	return op(a, b)
}

func sum(args ...int) int {
	result := 0
	for _, v := range args {
		result += v
	}
	return result
}

func main() {
	if result, err := eval(3, 4, "o"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(3, 4, "o"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(
		apply(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 2, 4),
	)

	fmt.Println(sum(1, 2, 3))
}
