package main

import "fmt"

func main() {

	var myGreetings map[string]string
	// myGreetings["Tim"] = "Good morning."
	// myGreetings["Jenny"] = "Bonjour."

	fmt.Println(myGreetings)
	fmt.Println(myGreetings == nil)
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
