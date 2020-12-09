package main

import "fmt"

func main() {
	fmt.Print(
		grade(0),
		grade(50),
		grade(60),
		grade(70),
		grade(80),
		grade(90),
		grade(100),
		//grade(110),
	)
}

func grade(score int) string {
	result := ""
	switch {
	case score <= 60:
		result = "E"
	case score <= 70:
		result = "D"
	case score <= 80:
		result = "C"
	case score <= 90:
		result = "B"
	case score <= 100:
		result = "A"
	default:
		panic("找不到匹配值")
	}
	return result
}
