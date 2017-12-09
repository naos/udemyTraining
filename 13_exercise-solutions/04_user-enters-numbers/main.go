package main

import "fmt"

func main() {
	var smallnum int
	var largenum int

	fmt.Print("Plase enter a small number:")
	fmt.Scan(&smallnum)

	fmt.Print("Please enter a large number:")
	fmt.Scan(&largenum)

	remnum := largenum % smallnum
	fmt.Printf("The reminder of %d divided by %d is %d\n", largenum, smallnum, remnum)
}
