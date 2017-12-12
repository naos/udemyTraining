package main

import "fmt"

func main() {

	myGreetings := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreetings)
	delete(myGreetings, "two")
	fmt.Println(myGreetings)
}
