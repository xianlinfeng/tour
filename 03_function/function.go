package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

// return value is signed a name(x, y), which is seen as declare variable at the top of func
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Printf("the sum of 3 and 4 is %v \n", add(3., 4))
	a, b := swap(3, 4)
	fmt.Printf("the swap of 3 and 4 is %v and %v \n", a, b)
	a, b = split(12)
	fmt.Printf("the split of 12 is %v and %v \n", a, b)

}
