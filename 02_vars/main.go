package main

import (
	"fmt"
)

func main() {
	// var name string = "Artur"
	// var age int = 22
	// var name = "Artur"
	// var age int32 = 22
	const isCool = true

	//Shorthand
	// name := "Alona"
	// size := 1.3

	name, email := "Artur", "artur@mail.com"

	fmt.Println(name, email, isCool)
	fmt.Printf("%T\n", email)
}
