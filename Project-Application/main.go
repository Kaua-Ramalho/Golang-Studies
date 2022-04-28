package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Initialize")

	application := app.Create()
	error := application.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}
