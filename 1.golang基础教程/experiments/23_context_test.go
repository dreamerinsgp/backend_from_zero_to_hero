package experiments

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// ctx, cancel := context.WithCancel(context.Background())

	// go func() {
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("context canceled")
	// 	case <-time.After(2 * time.Second):
	// 		fmt.Println("timeout")
	// 	}
	// }()

	// cancel()

	// time.Sleep(time.Second)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		for range 20 {
			select {
			case <-ctx.Done():
				fmt.Println("context canceled")
			case <-time.After(2 * time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	time.Sleep(6 * time.Second)

}
