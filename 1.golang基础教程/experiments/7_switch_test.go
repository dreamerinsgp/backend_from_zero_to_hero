package experiments

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	var i = 2
	switch i {
	case 1:
		fmt.Println("i 值为1")
	case 2:
		fmt.Println("i 值为2")
	default:
		fmt.Println("i 值为0")
	}

	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afater noon")
	}

}
