package main

import "fmt"

func main() {
	fmt.Println("Internal Arrays")

	// If you want to do a Slice, but you want to determine how many positions need to have, you can do it:
	slice := make([]int8, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice)) // You can use "len()" to see how much positions any slice has already filled
	fmt.Println(cap(slice)) // You can use "cap()" to see what's the total capacity of a Slice, without it's filled positions

	// This slice was created with ten positions filled with numbers, but it has totally fifteen positions, with or without filled it.
	// This is how a slice is created without fixed parameters, ir without an Array.

	// Even if you have crossed the final position, slices have dynamic allocation, so it will duplicate its limit. Example:
	slice = append(slice, 6)
	slice = append(slice, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// It is how a Slice works.
}
