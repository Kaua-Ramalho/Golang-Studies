package main

import (
	"fmt"
	"modulo/testando"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá, Mundo!")
	testando.Escrever()

	erro := checkmail.ValidateFormat("hellothere@gmail.com")
	fmt.Println(erro)
}
