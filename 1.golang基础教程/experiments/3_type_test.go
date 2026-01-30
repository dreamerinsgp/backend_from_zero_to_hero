package experiments

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	//1. 使用type 来声明一个结构体类型
	type Person struct {
		age  int
		name string
	}

	var person Person

	fmt.Println(person)

	//2. 使用type 来声明一个接口
	type Animal interface {
		genre() string
		color() string
	}
}
