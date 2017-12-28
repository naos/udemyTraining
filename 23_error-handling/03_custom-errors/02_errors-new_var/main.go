package main

import (
	"errors"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	// implementation
	return 42, nil
}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
