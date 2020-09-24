package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("The sum of 0 to 100 is ", sum)

	sum = 1
	for sum < 1000 { // for can only have one content
		sum += sum
	}
	fmt.Println("The result of sum += sum is ", sum)

	if a := 3; a < 10 {
		fmt.Println("if statement 1")
	} else if a < 20 {
		fmt.Println("if statement 2")
	} else {
		fmt.Println("if statement 3")
	}
}
