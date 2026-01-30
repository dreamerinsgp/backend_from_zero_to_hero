package experiments

import (
	"fmt"
	"math"
	"testing"
)

type geometry interface {
	area() int
	perim() int
}

type rectangle struct {
	width, height int
}

type circle struct {
	radius int
}

func (r rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func (c circle) area() int {
	return int(math.Pi * float64(c.radius) * float64(c.radius))
}

func (c circle) perim() int {
	return int(2 * math.Pi * float64(c.radius))
}

func TestInterface(t *testing.T) {
	r := rectangle{3, 5}
	c := circle{2}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	fmt.Println(c.area())
	fmt.Println(c.perim())
}
