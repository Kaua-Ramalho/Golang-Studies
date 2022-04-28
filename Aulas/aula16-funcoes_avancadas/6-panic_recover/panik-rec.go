package main

import (
	"fmt"
)

func recuperarFuncao() {
	if r := recover(); r != nil {
		fmt.Println("Function successfully recovered!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarFuncao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!!!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Post Processing")
}
