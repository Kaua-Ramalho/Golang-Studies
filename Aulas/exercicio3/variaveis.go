package main

import "fmt"

func main() {
	var variavel1 string = "Hello, World!"
	var variavel2 int = 1234
	fmt.Println(variavel1, variavel2)

	var (
		variavel3 = "Variáveis"
		variavel4 = "Várias variáveis"
	)
	fmt.Println(variavel3, variavel4)

	variavel5 := "Alalal"
	variavel6 := "Lalala"

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
