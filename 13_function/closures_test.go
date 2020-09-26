package main

import (
	"fmt"
	"testing"
)

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestClouses(t *testing.T) {
	pos := add()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
