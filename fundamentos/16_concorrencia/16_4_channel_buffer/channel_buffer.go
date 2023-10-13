package main

import "fmt"

func main() {

	// Nao funciona
	// Essa ação gera um deadlock
	// pois é uma ação bloqueante

	/*
		fmt.Println("Channel Buffer")

		canal := make(chan string)

		canal <- "Olá Mundo!" // Envia o texto para o canal (bloqueante)

		mensagem := <-canal // Aguarda o canal ser preenchido

		fmt.Println(mensagem)
	*/

	// Canal com buffer
	// canal := make(chan string, 1) // Canal com buffer de 1 posição
	// Só bloqueia quando o canal está cheio
	canal := make(chan string, 2) // Canal com buffer de 2 posições

	canal <- "Olá Mundo!" // Envia o texto para o canal

	mensagem := <-canal // Aguarda o canal ser preenchido

	fmt.Println(mensagem)
}
