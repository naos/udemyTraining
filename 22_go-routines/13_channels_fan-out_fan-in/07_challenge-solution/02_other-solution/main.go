package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println(n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel
-- Use the "fan out / fan in" pattern to accomplish this

CHALLENGE #2:
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
