package main

import (
	"fmt"
	"modulo/tipos"
)

func main() {
	// Em Go nós temos todos os Dados Primitivos que conhecemos, no entanto, eles têm suas subvertentes, que são específicadas
	var (
		numero1 int8  = 123   // int8 é usada para especificar que o número só pode ter até 8 bits de tamanho
		numero2 int16 = -1234 // int16 é usada para especificar que o número só pode ter até 16 bits de tamanho
		numero3 int32 = 1256  // int32 é usada para especificar que o número só pode ter até 32 bits de tamanho
		numero4 int64 = -1094 // int64 é usada para especificar que o número só pode ter até 64 bits de tamanho
		numero5 int   = 1246  // int é usada de acordo com a arquitetura do computador. Ou seja: Se ele for 64 bits, int irá usar 64 bits
	)
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)

	// Há, ainda uma subvariação de int, que seria a uint, que significa "unsigned int". Ele não pode conter sinal de negativo.
	var unumero uint = 1234 // uint segue a mesma lógica de bits do int
	fmt.Println(unumero)

	// Floats: Seguem a mesma lógica de bits dos int, com diferença que aceitam apenas 32 e 64 bits ou a arquitetura da máquina
	var (
		flutuante1 float32 = 1.757
		flutuante2 float64 = 123.7908709759028734
	)
	fmt.Println(flutuante1)
	fmt.Println(flutuante2)

	fmt.Println()
	fmt.Println()

	tipos.Escrever()
}
