package operator_test

import "testing"

// 用 == 比较数组
// 1. 相同维数且含有相同个数元素的数组才可以比较
// 2. 每个元素都相同的才相等

// 位运算符
// &^ 按位置零
// 1 &^ 0 -- 1
// 1 &^ 1 -- 0
// 0 &^ 1 -- 0
// 0 &^ 0 -- 0
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
}
