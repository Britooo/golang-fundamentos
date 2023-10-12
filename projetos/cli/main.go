package main

import (
	"cli-ip-discover/app"
	"log"
	"os"
)

func main() {
	//Caso tenha erro com o gopls, abra esse modulo no vscode para parar de dar erro
	// Ou habilita o go modules

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
