package main

import (
	"fmt"
	"time"
)

/* type error interface {
	Error() error
}
*/

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v : %s ", e.When, e.What)
}

func run() error {
	return MyError{time.Now(), "I dont know."}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
