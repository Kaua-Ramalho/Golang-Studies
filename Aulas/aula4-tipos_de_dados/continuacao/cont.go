package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Há também 'char', que é usada em muitas outras linguagens para conversão de variáveis. " +
		"No entanto, aqui essa função não é usada dessa forma. Ela apenas retorna o valor da tabela ask")
	fmt.Println("Exemplo:")

	char := 'B'
	fmt.Println(char)
	fmt.Println("Foi usada a função 'char' aderindo ela ao valor 'B', porém ela foi retornada com o número 66.")

	fmt.Println("Detalhe: Não é possível usar aspas duplas e nem mais de um caractere no char. Exemplo:")
	fmt.Println("Não é possível usar char := \"B\" ou char := 'T3'. Apenas char :='(um caractere só)'")
	fmt.Println()

	// Boolean e "Erros"
	fmt.Println("Em Go, nós também temos os booleanos e 'erros', que também são tipos, pois a linguagem dá um tratamento diferente para erros." +
		"Exemplo:")
	var booleano bool = false
	fmt.Println(booleano)

	fmt.Println()

	fmt.Println("Erros servem de valor zero para muitos elementos nos quais não conseguem ter valor zero, " +
		"como ponteiros, os próprios erros e interfaces. Exemplo:")

	var erro error
	fmt.Println(erro)
	fmt.Println()
	fmt.Println("Como visto, o valor retornado é <nil>. Esse valor funciona como 'null' ou 'nulo'. " +
		"Há possibilidade de inserir valores para esse tipo de dado, porém é necessário importar o pacote 'errors' e " +
		"usar uma função própria. Exemplo: ")

	var errado error = errors.New("Deu erro aqui, cara. Presta atenção, bicho burro! ")
	fmt.Println(errado)
}
