package main

import (
	"fmt"
)

// Essa é a primeira forma de declarar um switch.
func diaDaSemana(numero int16) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// Há outra forma de declarar o Switch, também.
func diaDaSemana2(numero2 int8) string {
	switch {
	case numero2 == 1:
		return "Domingo"
	case numero2 == 2:
		return "Segunda-Feira"
	case numero2 == 3:
		return "Terça-Feira"
	case numero2 == 4:
		return "Quarta-Feira"
	case numero2 == 5:
		return "Quinta-Feira"
	case numero2 == 6:
		return "Sexta-Feira"
	case numero2 == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func main() {
	fmt.Println("Switches")

	var dia string = diaDaSemana(7)
	fmt.Println(dia)

	var dia2 string = diaDaSemana2(3)
	fmt.Println(dia2)
}
