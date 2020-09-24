package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "No"
	k := 3
	fmt.Println(c, python, java, k)

	var a int
	fmt.Println("The zero value of int is ", a)
}

// type conversoin
func varChangeType() uint {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	return u
}
