package main

import "fmt"

func main() {
	fmt.Println(greatest(2, 4, 78, 9, 3, 1))
}

func greatest(xs ...int) int {
	var mn int
	for _, n := range xs {
		if n > mn {
			mn = n
		}
	}
	return mn
}
