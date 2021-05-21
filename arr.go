package main

import "fmt"

func doArr1(arr *[3]int) {
	arr[0] = 2
	for _, v := range arr {
		fmt.Println(v)
	}
}

func doArr2(arr []int) {
	arr[0] = 100
}

func main() {
	fmt.Println("arr")
	var arr1 [3]int
	fmt.Println(arr1)
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	arr3 := [...]int{2, 4, 6, 8}
	fmt.Println(arr3)
	fmt.Println("doArr1(arr1)")
	doArr1(&arr1)
	fmt.Println(arr1)

	arr4 := [3]int{1, 2, 3}
	s1 := arr4[1:2]
	fmt.Println("cap(s1)", s1)
	s2 := append(s1, 10)
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("arr4", arr4)

	arr5 := [3]int{1, 2, 3}
	doArr2(arr5[:])
	fmt.Println("arr5", arr5)

	arr6 := []int{1, 2, 3, 4, 5, 6}
	s3 := append(arr6[1:2], arr6[1:4]...)
	fmt.Println("s3", s3)
	fmt.Println("cap(s3)", cap(s3))
	fmt.Println("len(s3)", len(s3))
	fmt.Println("cap(arr6)", cap(arr6))
	fmt.Println("len(arr6)", len(arr6))
}
