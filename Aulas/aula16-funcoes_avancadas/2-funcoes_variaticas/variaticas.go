package main

import "fmt"

func soma(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func print1(text string, numbers2 ...int) {
	for _, number2 := range numbers2 {
		fmt.Println(text, number2)
	}
}
func main() {
	totalDaSoma := soma(10, 400, 10000000)
	fmt.Println(totalDaSoma)

	print1("Hello, World", 1, 2, 3, 4, 5, 6)
}
