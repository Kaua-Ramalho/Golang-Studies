package main

import (
	"fmt"
)

// Há duas formas de declarar funções, em Go. A primeira forma é essa:
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
} // Essa forma acaba sendo igual à declaração da função main.

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// E a segunda forma é essa:
	var f = func(text string) string {
		return text
	} // Essa forma é a mesma coisa de declarar uma variável, porém em forma de função. Isso acaba sendo útil para economizar linhas.

	retorno := f("Hello, World")
	fmt.Println(retorno)
}
