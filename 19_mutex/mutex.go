package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter a safe counter
type SafeCounter struct {
	m map[string]int
	t sync.Mutex
}

func (c *SafeCounter) Count(key string) {
	c.t.Lock()
	c.m[key]++
	c.t.Unlock()
}

func (c *SafeCounter) Increase(key string) {
	c.m[key]++
}

func (c *SafeCounter) Value(key string) (i int) {
	c.t.Lock()
	i = c.m[key]
	c.t.Unlock()
	return
}

func main() {
	c := &SafeCounter{m: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		go c.Count("Somekey")
	}

	time.Sleep(time.Second * 3)
	fmt.Println(c.Value("Somekey"))

	d := &SafeCounter{m: make(map[string]int)}
	fmt.Println("Count without safety lock:")
	for i := 0; i < 30; i++ {
		go d.Increase("Somekey")
	}

	time.Sleep(time.Second * 3)

	fmt.Println(d.m["Somekey"])
}
