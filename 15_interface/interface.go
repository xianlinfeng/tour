package main

import (
	"fmt"
	"math"
)

// Abser a ABS function
type Abser interface {
	Abs() float64
}

type Float float64

func (f Float) Abs() float64 {
	if f < 0 {
		return float64(-f)

	} else {
		return float64(f)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/

func main() {
	var a Abser
	f := Float(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v
	fmt.Println(a.Abs())

}
