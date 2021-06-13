package fib

import "testing"

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 2

	// var (
	// 	a = 1
	// 	b = 2
	// )

	a, b := 1, 2

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = a + tmp
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
