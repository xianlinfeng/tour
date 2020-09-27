package exercise

import (
	"io"
	"os"
	"strings"
	"testing"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b) // access to the(struct) element's method
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

func TestRotReader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
