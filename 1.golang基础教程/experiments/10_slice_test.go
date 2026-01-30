package experiments

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	//1.s初始默认值为nil ,长度为0
	var s []string
	fmt.Println(s == nil)
	fmt.Println(len(s) == 0)

	//2. 创建带指定len的slice
	var s1 = make([]string, 3)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	//3. 指定cap
	var s2 = make([]string, 3, 5)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	//4. append
	var s3 = make([]string, 3)
	s3[0] = "1"
	s3[1] = "2"
	s3[2] = "3"
	fmt.Println(s3)
	s4 := append(s3, "4")
	fmt.Println(s4)
}
