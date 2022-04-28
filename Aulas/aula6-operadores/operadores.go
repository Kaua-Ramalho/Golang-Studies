package main

import (
	"fmt"
)

func main() {
	// Operadores aritméticos
	soma := 20 + 40
	subtracao := 432 - 50
	multiplicacao := 42 * 2
	divisao := 50 / 5
	restoDaDivisao := 50 % 4
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// Operadores de Atribuição
	var texto string = "Testing"
	texto2 := "Teste"
	fmt.Println(texto, "\n", texto2)

	// Operadores Relacionais
	fmt.Println(1 > 3)
	fmt.Println(1 >= 0)
	fmt.Println(50 < 51) // Esses operadores serão booleanos
	fmt.Println(1 == 1)
	fmt.Println(1 != 2)

	fmt.Println("-----------------")

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro, !falso)

	fmt.Println("----------------")

	// Operadores Unários
	numero := 30
	numero++
	numero += 19 // São formas de somar, subtrair, multiplicar, dividir ou extrair o resto da divisão, sem ter que declarar a variável novamente com a expressão
	fmt.Println(numero)

	numero2 := 20
	numero2--

	numero3 := 3
	numero3 *= 3

	numero4 := 10
	numero4 /= 2

	fmt.Println(numero2, numero3, numero4)

	// Operadores Tenários não existem em Go. O jeito de replicar esses operadores são com if e else
}
