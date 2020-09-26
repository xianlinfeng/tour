package main

import (
	"fmt"
	"math"
)

// Vertex struct
type Vertex struct {
	X, Y float64
}

// Abs function
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale function
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// NewAbs New Abs function
func NewAbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// NewScale function
func NewScale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v)

	fmt.Println(NewAbs(v))
	NewScale(&v, 10)
	fmt.Println(v)
}
