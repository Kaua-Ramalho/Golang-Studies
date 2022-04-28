package main

import (
	"fmt"
)

type test struct {
	teste  string
	teste2 int16
}

func somar(num1, num2 int16) (int16, int16) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}

func main() {
	fmt.Println("Testando")

	retornoSoma, retornoSubtracao := somar(16, 30)
	fmt.Println(retornoSoma, "\n", retornoSubtracao)

	testing := test{"Ava", 18}
	fmt.Println(testing)

	ola := [50]int8{1, 2, 4, 5, 6, 7}
	fmt.Println(ola)
}
