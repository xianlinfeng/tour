package exercise

import (
	"fmt"
	"testing"
)

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}

func TestClosure(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
