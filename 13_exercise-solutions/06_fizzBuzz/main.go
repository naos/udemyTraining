package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		tag := ""
		if i%3 == 0 {
			tag += "Fizz"
		}
		if i%5 == 0 {
			tag += "Buzz"
		}
		if tag != "" {
			fmt.Println(tag)
		} else {
			fmt.Println(i)
		}
		// fmt.Printf("%d %s\n", i, tag)
	}
}
