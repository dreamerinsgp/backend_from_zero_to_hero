package experiments

import (
	"fmt"
	"sync"
	"testing"
)

type Container struct {
	mu    sync.Mutex
	count map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count[name]++
}

func TestMutex(t *testing.T) {
	var c = Container{
		count: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	var update = func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Go(func() {
		update("a", 10000)
	})

	wg.Go(func() {
		update("a", 10000)
	})

	wg.Go(func() {
		update("b", 200)
	})

	wg.Wait()

	fmt.Println(c.count["a"], c.count["b"])
}
