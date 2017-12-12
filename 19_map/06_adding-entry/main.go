package main

import "fmt"

func main() {

	myGreetings := map[string]string{
		"Tim":   "Good morning.",
		"Jenny": "Bonjour.",
	}

	myGreetings["Harleen"] = "Howdy"

	fmt.Println(myGreetings)
}
