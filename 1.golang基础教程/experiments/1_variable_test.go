package experiments

import (
	"fmt"
	"testing"
)

var a = "initial"

func TestVariable(t *testing.T) {
	fmt.Println(a)

	f := "apple"
	fmt.Println(f)
}
