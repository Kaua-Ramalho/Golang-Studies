// Criação de modulos é mais útil quando há mais de um pacote para se manusear. Apesar disso, a extensão do VS acaba pedindo módulos.
package main

import (
	"fmt"
)

func main() {
	// Há duas formas de declarar variáveis. Dessa forma:
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// E dessa forma:
	variavel2 := "Hello, World!" //Detalhe: O VS não deixa você inutilizar uma variável criada. Criou? Use!
	fmt.Println(variavel2)

	// Há subvariações de declaração, também, que economizam linhas de código. Exemplos:
	var (
		variavel3 string  = "Variável 3"
		variavel4 float64 = 1.59
	)
	variavel5, variavel6 := "Variável 5", 43 // Nesse caso, os valores são separados por vírgula e seguem a ordem declarada

	fmt.Println(variavel3, variavel4, variavel5, variavel6)
}
