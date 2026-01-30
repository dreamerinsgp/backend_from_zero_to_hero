package experiments

import (
	"fmt"
	"testing"
)

func outer() func() int {
	x := 0 // 外部变量
	return func() int {
		x++ // 闭包捕获并修改外部变量
		return x
	}
}
func TestClosure(t *testing.T) {

	f := outer()     // outer函数返回后，x仍然存在
	fmt.Println(f()) // 输出: 1
	fmt.Println(f()) // 输出: 2
	fmt.Println(f()) // 输出: 3
	fmt.Println(f())
}
