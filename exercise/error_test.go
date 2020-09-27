package exercise

import (
	"fmt"
	"testing"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number <%f> \n", e)
}

func SqrtErr(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func TestError(t *testing.T) {
	fmt.Println(SqrtErr(2))
	fmt.Println(SqrtErr(-2))
}
