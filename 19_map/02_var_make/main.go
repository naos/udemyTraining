package main

import "fmt"

func main() {

	var myGreetings = make(map[string]string)
	myGreetings["Tim"] = "Good morning."
	myGreetings["Jenny"] = "Bonjour."

	fmt.Println(myGreetings)
}
