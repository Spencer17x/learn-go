package string_test

import "testing"

// 于其他主要编程语言的差异
// 1. string 是数据类型，不是引用或指针类型
// 2. string 是只读的 byte slice，len 函数可以它所包含的 byte 数
// 3. string 的 byte 数组可以存放任何数据

// Unicode UTF8
// 1. Unicode 是一种字符集（code point）
// 2. UTF8 是 unicode 的存储实现（转换为字节序列的规则）
func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认零值""
	s = "hello"
	t.Log(len(s))
	// s[1] = "3" // string 是不可变的 byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	t.Log(len(s)) // 是 byte 数
	s = "中"
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
