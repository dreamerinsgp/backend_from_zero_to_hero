package experiments

import (
	"fmt"
	"testing"
	"time"
)

func printlnHelloWorld() {
	fmt.Println("Hello,World!")
}

func TestGoroutine(t *testing.T) {

	go printlnHelloWorld()

	go func() {
		fmt.Println("Hello again!")
	}()

	time.Sleep(time.Second)
}
