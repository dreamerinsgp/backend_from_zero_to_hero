package experiments

import (
	"fmt"
	"testing"
)

// 1.结构体的声明
type person struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	var p person
	fmt.Println(p)
	p = person{
		name: "wang",
		age:  30,
	}
	fmt.Println(p)

	p.name = "chen"
	p.age = 20

	fmt.Println(p)
}
