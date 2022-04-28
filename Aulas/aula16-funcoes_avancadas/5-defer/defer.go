package main

import "fmt"

func function() {
	fmt.Println("Function 1")
}

func function2() {
	fmt.Println("Function 2")
}

func aprovacaoDeAluno(n1, n2 float32) bool {
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	defer function() // DEFER = ADIAR
	function2()
	nota := aprovacaoDeAluno(5.5, 8.0)
	fmt.Println(nota)
}
