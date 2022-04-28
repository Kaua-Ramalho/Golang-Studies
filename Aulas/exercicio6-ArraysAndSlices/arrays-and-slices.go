package main

import "fmt"

func main() {
	fmt.Println("Exercising")

	array1 := [3]int{1, 2, 3}
	fmt.Println(array1)

	var array2 [6]int8
	array2[0] = 2
	array2[1] = 4
	array2[2] = 6
	array2[3] = 8
	array2[4] = 10
	array2[5] = 12

	fmt.Println(array2)

	fmt.Println("--------------------------------------------------------------------------")

	slice1 := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// If you want to fill a position in one slice, you can do it:
	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	slice2 := array2[0:5]
	fmt.Println(slice2)
}
