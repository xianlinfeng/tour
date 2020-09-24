package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My faverite number is", rand.Intn(10))
	fmt.Printf("the sqrt of 2 is %v\n", math.Sqrt(2))
	fmt.Printf("the Pi is %v\n", math.Pi)
}
