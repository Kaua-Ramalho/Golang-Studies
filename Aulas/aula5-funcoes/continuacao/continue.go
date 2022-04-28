package main

import "fmt"

func calculo(n1, n2 int16) (int16, int16) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultadoSoma, resultadoSubtracao := calculo(10, 50)
	fmt.Println(resultadoSoma, "\n", resultadoSubtracao)
}
