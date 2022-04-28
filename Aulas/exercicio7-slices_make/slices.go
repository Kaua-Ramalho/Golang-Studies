package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercising")

	slice1 := make([]int8, 40)
	slice1 = append(slice1, 20)
	slice1 = append(slice1, 22)
	slice1 = append(slice1, 24)
	slice1 = append(slice1, 26)
	slice1 = append(slice1, 28)
	slice1 = append(slice1, 30)
	slice1 = append(slice1, 32)
	slice1 = append(slice1, 34)
	fmt.Println(slice1)

	fmt.Println("-------------------------")

	array1 := [8]int8{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array1)

	array1[7] = 10
	fmt.Println(array1)

	slice2 := array1[0:8]
	fmt.Println(slice2)
}
