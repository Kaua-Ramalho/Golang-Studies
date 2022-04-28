package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int16) (soma int16, subtracao int16) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
func main() {
	fmt.Println("Advanced Functions 1")

	soma, subtracao := calculosMatematicos(1000, 800)
	fmt.Println(soma, "\n", subtracao)
}
