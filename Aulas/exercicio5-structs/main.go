package main

import "fmt"

type endereco struct {
	rua    string
	numero uint16
}

type usuario struct {
	nome     string
	idade    uint8
	altura   float32
	endereco endereco
}

func main() {
	enderecoDoUsuario := endereco{"Travessa Paranamirim", 70}

	usuario2 := usuario{"Kauã Ramalho Leandro", 17, 1.86, enderecoDoUsuario}
	fmt.Println(usuario2)
}
