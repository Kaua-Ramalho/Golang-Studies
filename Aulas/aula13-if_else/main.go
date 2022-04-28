package main

import (
	"fmt"
)

func main() {
	fmt.Println("If and Else")

	var number int8 = 18
	if number > 15 {
		fmt.Println("The number is bigger than fifteen")
	} else {
		fmt.Println("The number is smaller or equal than 15")
	}

	if otherNumber := number; otherNumber > 16 { // We can't use variables created to if and else out of these structures
		fmt.Println("This number is bigger than sixteen")
	} else {
		fmt.Println("This number is smaller or equal than sixteen")
	}
}
