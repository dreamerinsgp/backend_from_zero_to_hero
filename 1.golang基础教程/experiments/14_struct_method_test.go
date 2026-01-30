package experiments

import (
	"fmt"
	"testing"
)

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestStructMethod(t *testing.T) {
	r := rect{width: 10, height: 10}

	area := r.area()

	fmt.Println(area)

	r1 := &r
	perim := r1.perim()
	fmt.Println(perim)
}
