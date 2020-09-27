package exercise

import (
	"fmt"
	"io"
	"testing"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	l := len(b)
	for i := range b {
		b[i] = 65
	}
	return l, nil
}

func TestReader(t *testing.T) {
	// r := strings.NewReader("Hello, Reader!用数据填充给定的字节切片并返回填充的字节数和错误值")
	r := MyReader{}
	b := make([]byte, 8)

	for i := 0; i < 10; i++ {
		n, err := r.Read(b)
		fmt.Printf("n = %v  err = %v b = %v \n ", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
