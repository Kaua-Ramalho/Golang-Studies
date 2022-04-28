package main

import "fmt"

func main() {
	return1 := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Parameters")

	fmt.Println(return1)
}
