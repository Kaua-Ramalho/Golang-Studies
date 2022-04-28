package main

import "fmt"

func main() {
	fmt.Println("Infernal pointers")

	var variable int
	var variable2 int

	variable = 10
	variable2 = variable

	variable += 5
	fmt.Println(variable, variable2) // Pointers are copies of variables, but without post changes.

	fmt.Println("----------------")

	// Pointers can also be memory references, on your code. Example:
	var variable3 int32
	var pointer *int32

	variable3 = 25
	pointer = &variable3

	fmt.Println(variable3, pointer)

	variable3 = 50
	fmt.Println(variable3, pointer)

}
