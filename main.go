package main

import (
	"cli-demo/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Kick off")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
