package experiments

import (
	"fmt"
	"testing"
)

const s = "web3"

func TestConst(t *testing.T) {
	fmt.Println(s)

	const n = 50000
	fmt.Println(n)

	const b = true
	fmt.Println(b)
}
