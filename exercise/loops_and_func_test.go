package exercise

import (
	"fmt"
	"testing"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func TestSqrt(t *testing.T) {
	t.Log(Sqrt(7))
}
