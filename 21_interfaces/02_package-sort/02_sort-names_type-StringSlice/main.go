package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)
	sort.Sort(sort.StringSlice(s))
	// sort.StringSlice(s).Sort()
	fmt.Println(s)
}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
