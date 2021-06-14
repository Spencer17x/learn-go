package type_test

import "testing"

type MyInt int64

// 类型转化
// 1. Go语言不允许隐式类型转换
// 2. 别名和原有类型也不能进行隐式类型转换

// 类型的预定义值
// 1. math.MaxInt64
// 2. math.MaxFloat64
// 3. math.MaxUint32

// 指针类型
// 1. 不支持指针运算
// 2. string 是值类型，其默认的初始值为空字符串，而不是 nil
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // 输出类型
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
