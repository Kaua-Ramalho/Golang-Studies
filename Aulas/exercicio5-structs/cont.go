package main

import "fmt"

type enderecoDeUsuario struct {
	rua    string
	numero uint16
}

type nomeDeUsuario struct {
	nome              string
	idade             uint16
	enderecoDeUsuario enderecoDeUsuario
}

func main() {
	endereco := enderecoDeUsuario{"Travessa Paranamirim", 70}
	usuario := nomeDeUsuario{"Kauã Ramalho Leandro", 17, endereco}
	fmt.Println(usuario)
}
