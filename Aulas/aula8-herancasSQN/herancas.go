package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pessoa1 := pessoa{"Kauã", "Ramalho Leandro", 17, 1.86}
	estudante1 := estudante{pessoa1, "Engenharia da Computação", "Eniac"}

	fmt.Println(estudante1)
}
