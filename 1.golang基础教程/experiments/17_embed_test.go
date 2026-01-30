package experiments

import (
	"fmt"
	"testing"
)

type A struct {
	a1, a2 int
	B
}

type B struct {
	b1, b2 int
}

func (b B) method1() int {
	return b.b1 + b.b2
}

func TestEmbed(t *testing.T) {
	b := B{10, 20}

	a := A{1, 2, b}

	b1 := a.b1
	b2 := a.b2
	sum := a.method1()
	fmt.Println(b1, b2, sum)
}
