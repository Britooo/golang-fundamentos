package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {

		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	/*
		for {
			mensagem1 := <-canal1 // deveria esperar meio segundo, mas como o canal2 demora 2 segundos, o canal1 fica esperando
			println("Mensagem 1:", mensagem1)

			mensagem2 := <-canal2 // espera dois segundos
			println("Mensagem 2:", mensagem2)
		}
	*/

	fmt.Println("Select")

	// Select é uma estrutura de controle que permite executar uma ação que estiver pronta primeiro
	// Se o canal1 estiver pronto, ele executa o case do canal1
	// Se o canal2 estiver pronto, ele executa o case do canal2
	// Serve para sincronizar as goroutines

	for {
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}
}
