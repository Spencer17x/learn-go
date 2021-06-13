package try_test

import "testing"

// 编写测试程序
// 1. 源码文件以 _test 结尾：xxx_test.go
// 2. 测试方法名以 Test 开头：func TestXXX(t *testing.T) { ... }

// 变量赋值
// 1. 赋值可以进行自动类型推断
// 2. 在一个赋值语句中可以对多个变量进行同时赋值
func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}
