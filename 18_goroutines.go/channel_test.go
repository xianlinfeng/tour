package main

import (
	"fmt"
	"testing"
)

func TestChannelWithBuffer(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func Fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		c := x
		x = y
		y += c
	}
	close(c)
}

func TestCloseChannel(t *testing.T) {
	ch := make(chan int, 10)
	go Fibonacci(15, ch)
	for v := range ch {
		fmt.Println(v)
	}

}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // random select a branch to excute
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		default:
			fmt.Println("The default branch to excute")
		}

	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
