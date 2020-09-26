package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// the interface of stringer
/* type Stringer interface {
    String() string
} */

func (p *Person) string() string {
	return fmt.Sprintf("%v : %v ", p.Name, p.Age)
}

func TestStringer(t *testing.T) {
	a := Person{"Nash", 31}
	b := Person{"Arthur", 28}

	fmt.Println(a.string())
	fmt.Println(b.string())

}
