package main

import (
	"fmt"
)

func invertSignal(number int) int {
	return number * -1
}

func newInvertSignal(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("Functions plus Pointers")
	number := 10
	newNumber := invertSignal(number)
	fmt.Println(newNumber)

	fmt.Println()

	number1 := 20
	fmt.Println(number1)
	newInvertSignal(&number1)
	fmt.Println(number1)
}
