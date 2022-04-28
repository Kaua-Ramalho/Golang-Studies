package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")

	// Arrays are values' lists.

	// There's two ways to use Arrays. This:
	var array1 [5]int32
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	fmt.Println(array1)

	// And this:

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// If you want to do a dynamic array, without fixed positions number, you need to do it:
	array3 := [...]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5", "Posição 6", "Posição 7"}
	fmt.Println(array3)
	// If you do it, you don't need to worry about limit of positions, but you cannot declare after curly brackets

	fmt.Println("-------------------------------------------------")

	// Slices are Arrays, but completely dynamics, without previous problems we've seen before.
	// If you want to use a Slice, you can use the same ways you used to use an Array, but there's the difference:
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 60, 55, 43}
	fmt.Println(slice)

	// If you want to add a value, without change anything on original slice, you can use "append" to do it. Example:
	slice = append(slice, 20)
	fmt.Println(slice)

	// You can use a slice of an Array, using Slice, too. Example:
	slice2 := array3[0:3]
	fmt.Println(slice2) // It's like a pointer
}
