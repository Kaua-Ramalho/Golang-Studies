package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":      "Kauã",
		"last name": "Ramalho Leandro",
	}
	fmt.Println(user)

	user2 := map[string]map[string]string{
		"Nome Completo": {
			"nome":      "Kauã",
			"sobrenome": "Ramalho Leandro",
		},
		"Faculdade": {
			"curso":       "Engenharia da Computação",
			"instituição": "ENIAC",
			"campus":      "Guarulhos",
		},
	}
	fmt.Println(user2)
}
