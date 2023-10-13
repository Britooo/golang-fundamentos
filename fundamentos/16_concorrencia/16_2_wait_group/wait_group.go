package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {

	fmt.Println("Wait Group")

	// WaitGroup: Espera um grupo de goroutines terminarem
	// Especie de await (JS)
	var waitgroup sync.WaitGroup

	// Adiciona quantas goroutines serão executadas
	waitgroup.Add(2)

	go func() {
		escrever("Goroutine 1")
		waitgroup.Done() // -1 sinaliza que terminou
	}()

	go func() {
		escrever("Goroutine 2")
		waitgroup.Done() // -1 sinaliza que terminou
	}()

	// Espera as duas funções terminarem
	// Caso não seja chamado, o programa termina antes de executar as funções
	// Comente a linha abaixo e veja o que acontece
	waitgroup.Wait()
}
