package experiments

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	if 7%2 == 0 {
		fmt.Println("7 是偶数")
	} else {
		fmt.Println("7 是奇数")
	}
}
