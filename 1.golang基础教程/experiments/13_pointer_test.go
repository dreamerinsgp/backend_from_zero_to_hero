package experiments

import (
	"fmt"
	"testing"
)

func Pointer(ptr *int) {
	*ptr = 2
}

func Value(val int) {
	val = 2
}

func TestPointer(t *testing.T) {
	//1.传递指针
	i := 1
	fmt.Println(i)

	Pointer(&i)

	fmt.Println(i)

	//比较指针传递和值传递的区别？
	k := 1
	fmt.Println(k)

	Value(k)
	fmt.Println(k)
}
