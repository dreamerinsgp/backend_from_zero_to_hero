package experiments

import (
	"testing"
)

func TestChan(t *testing.T) {
	//1.无缓冲区chan
	// ch := make(chan int)

	// go func() {
	// 	val := <-ch
	// 	fmt.Println(val)
	// }()

	// ch <- 10

	//2.有缓冲区
	ch1 := make(chan int, 10)

	ch1 <- 10
}
