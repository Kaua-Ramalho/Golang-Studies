package main

import (
	"fmt"
)

func calcular(n1, n2 int32) (int32, int32) {
	somar := n1 + n2
	subtrair := n1 - n2
	return somar, subtrair
}

func main() {
	resultadoSoma, resultadoSubtrair := calcular(4321, 8765)
	fmt.Println(resultadoSoma, "\n", resultadoSubtrair)

	fmt.Println()

	var escrita = func(text string) string {
		return text
	}
	escrever := escrita("Escrevendo de uma variável função")
	fmt.Println(escrever)
}
