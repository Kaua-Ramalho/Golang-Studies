package main

import "fmt"

func generic(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generic("AA")
	generic(1234)
	generic(1.25)
	generic(true)

	fmt.Println("AA", 123, 1.25, true)
}
