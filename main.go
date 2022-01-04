package main

import (
	"fmt"
	"get-ip/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Generate()
	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
