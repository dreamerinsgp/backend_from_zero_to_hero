package experiments

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	//1.声明数组
	var a [5]int
	fmt.Println(a)

	//2.给数组中的指定下标对应的元素位置赋值
	a[4] = 100
	fmt.Println(a)

	//3.声明同时赋值
	// b := [5]int{1, 2, 3, 4, 5}
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//4.采用 ... 来不显式指定数组大小
	c := [...]int{1, 2, 3}
	fmt.Println(c)
	fmt.Println(len(c))
}
