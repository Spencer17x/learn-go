package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

// panic
// panic 用于不可以恢复的错误
// panic 退出前会执行 defer 指定的内容

// pnic vs os.Exit
// os.Exit 退出时不会调用 defer 指定的函数
// os.Exit 退出时不输出当前调用栈信息
func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	fmt.Println(errors.New("Something wrong"))

	// os.Exit(-1)
	// fmt.Println("End")
	// panic("error panic")
}
