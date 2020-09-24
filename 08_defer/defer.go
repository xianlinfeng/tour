package main

import "fmt"

// defer: first in, last out

func main() {
	fmt.Println("At the top of the function")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("At the bottom of the function")
}
