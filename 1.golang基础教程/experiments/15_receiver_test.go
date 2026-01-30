package experiments

import (
	"fmt"
	"testing"
)

type rect1 struct {
	width, height int
}

func (r rect) valueUpdate() {
	r.width = 1
	r.height = 2
}

func (r *rect) pointerUpdate() {
	r.width = 20
	r.height = 20
}

func TestReceiver(t *testing.T) {
	r := rect{width: 10, height: 10}
	fmt.Println(r)

	r.valueUpdate()

	fmt.Println(r)

	r.pointerUpdate()

	fmt.Println(r)
}
