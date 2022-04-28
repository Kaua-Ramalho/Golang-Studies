package main

import (
	"fmt"
)

type userid struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo de Structs")

	// Para deixar uma struct visível no código, existem duas formas. Essa:
	var user userid
	user.nome = "Kauã Ramalho"
	user.idade = 17
	fmt.Println(user)

	// E essa:
	user2 := userid{"Kauã Ramalho", 17} // Essa forma é bem mais fácil e econômica
	fmt.Println(user2)
}
