package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	sort.Sort(studyGroup)
	sort.Sort(sort.StringSlice(s))
	sort.Sort(sort.IntSlice(n))
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)

	sort.Sort(sort.Reverse(studyGroup))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)
}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
