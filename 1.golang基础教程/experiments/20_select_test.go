package experiments

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "message2"
	}()

	for {
		select {
		case message1 := <-ch1:
			fmt.Println("received message1:", message1)
		case message2 := <-ch2:
			fmt.Println("received message2:", message2)
		}
	}
}
