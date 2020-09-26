package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	prime := primes[1:4] // create a slices
	fmt.Println(prime)

	// struct slices and the capacity/length of a slice
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{11, false},
	}
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d, %v \n", len(s), cap(s), s)

	// nil slices
	var p []int
	fmt.Printf("len = %d, cap = %d, %v \n", len(p), cap(p), p)
	if p == nil {
		fmt.Println("nil")
	}

	// make a slice
	b := make([]int, 5)

	c := make([]int, 0, 5) // len(c) = 0, cap(c) = 5
	fmt.Println(b, c)

	// append a new element into a slice
	b = append(b, 3, 4, 5, 6)
	fmt.Println(b)

	// range
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range d {
		fmt.Println(i, ": ", v)
	}

	pow := make([]int, 32)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Println(value)
	}
}
