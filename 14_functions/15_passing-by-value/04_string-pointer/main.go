package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // 0xc4200140b8

	changeMe(&name)

	fmt.Println(&name) // 0xc4200140b8
	fmt.Println(name)  // Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0xc4200140b8
	fmt.Println(*z) // Todd
	*z = "Rocky"
	fmt.Println(z)  // 0xc4200140b8
	fmt.Println(*z) // Rocky
}
