package tipos

import "fmt"

func Escrever() {
	// Em GO nós também temos apelidos para os subtipos de dados. São chamados de Alias.

	var (
		// RUNE = INT32
		numeros rune = -12341234

		// BYTE = UINT8
		numeros2 byte = 255 // Para 8 bits é suportado até o número 255
	)
	fmt.Println(numeros)
	fmt.Println(numeros2)

}
