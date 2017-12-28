package main

import (
	"fmt"
	"runtime"
)

func main() {
	f := factorial(20)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	c := make(chan int)
	done := make(chan bool)
	numcpu := runtime.NumCPU()

	for i := 0; (i < numcpu) && (i < n); i++ {
		go func(i, n, step int) {
			subtotal := 1
			for j := n - i; j > 0; j -= step {
				subtotal *= j
				// fmt.Printf("goroutine %d number: %d\n", i, j)
			}
			c <- subtotal
			// fmt.Printf("goroutine %d subtotal: %d\n", i, subtotal)
			done <- true
		}(i, n, numcpu)
	}

	go func() {
		for i := 0; (i < numcpu) && (i < n); i++ {
			<-done
		}
		close(c)
		close(done)
	}()

	total := 1
	for n := range c {
		total *= n
	}
	return total
}
