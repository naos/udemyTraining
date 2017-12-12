package main

import "fmt"

func main() {

	myGreetings := make(map[string]string)
	myGreetings["Tim"] = "Good morning."
	myGreetings["Jenny"] = "Bonjour."

	fmt.Println(myGreetings)
}
