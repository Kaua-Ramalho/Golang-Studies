package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	// Fibonacci = 0 1 1 2 3 5 8 13 21 34 55...

	position := uint(6)
	fmt.Println(fibonacci(position))
}
