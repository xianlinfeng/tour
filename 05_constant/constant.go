package main

import (
	"fmt"
)

const Pi = 3.14
const World = "世界"
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println(float64(Big))
}
