package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	var i interface{} = "Hello "

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) // if i store a type of string, ok = ture and assign the value of i to s
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // 报错(panic)
	// fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's a int, and its value is %v \n", v)
	case string:
		fmt.Println("It's a string, the content is ", v)
	default:
		fmt.Printf("I dont know the type %T\n", v)
	}
}

func TestTypeSwitch(t *testing.T) {
	do(42)
	do("Hello")
	do(true)
}
