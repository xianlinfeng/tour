package main

import (
	"fmt"
	"math"
	"testing"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("%v : %T \n ", i, i) // interface has a tuple of avalue and a concrete type
}

// test of interface
func TestInterface(t *testing.T) {
	var i I

	i = &T{"Hello "}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

// nil interface
func TestNilInterface(t *testing.T) {
	var i I
	var tt *T
	i = tt // assign a nil value to a interface
	describe(i)
	i.M()

	var j I
	describe(j)
	// j.M()  nil interface cannot call a method (will preduce a error)
}

// empty interface
func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	describeInterface(i)

	i = 42
	describeInterface(i)

	i = "Hello"
	describeInterface(i)
}
