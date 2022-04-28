package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	var i int8 = 4
	for i < 5 {
		i++
		fmt.Println("Increasing i")
		time.Sleep(time.Nanosecond)
	}
	fmt.Println(i)

	for j := 0; j < 2; j++ {
		fmt.Println("Increasing j", j)
		time.Sleep(time.Nanosecond)
	}

	fmt.Println()

	names := []string{"Kauã", "Ramalho Leandro"}
	for indice, nome := range names {
		fmt.Println(indice, nome)
	}

	fmt.Println()

	maps := map[string]string{
		"name:":    "Kauã",
		"surname:": "Ramalho Leandro",
	}
	for chave, valor := range maps {
		fmt.Println(chave, valor)
	}

	fmt.Println()

	for indice2, letra := range "KAUÃ RAMALHO" {
		fmt.Println(indice2, string(letra))
	}

	// We can't create loops for Struct
}
