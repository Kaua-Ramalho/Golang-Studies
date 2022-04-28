package main

import (
	"fmt"
)

func closure() func() {
	text := "Inside closure's function"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	fmt.Println("Closure")

	text := "Inside main's function"
	fmt.Println(text)

	anotherFunction := closure()
	anotherFunction()
}
