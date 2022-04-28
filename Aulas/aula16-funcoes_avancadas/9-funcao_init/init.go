package main

import (
	"fmt"
)

func init() { // The Init is a function that will execute before main's function
	fmt.Println("Executando a função Init")
}

func main() {
	fmt.Println("Executing main function") // The difference between Main and Init is Init can be used in many archives inside a package
}
