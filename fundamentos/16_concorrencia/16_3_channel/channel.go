package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Envia o texto para o canal
		time.Sleep(time.Second)
	}

	close(canal) // Fecha o canal
}

func main() {

	fmt.Println("Channel")

	// Canal sem buffer
	// canal := make(chan string)
	// Serve para sincronizar as goroutines

	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// mensagem := <-canal // Aguarda o canal ser preenchido

	// O programa só termina quando o canal for preenchido
	// fmt.Println(mensagem)
	// O problema que na primeira mensagem o canal é preenchido e logo segue para o fim do programa

	// Porém, isso gera o famoso deadlock, pois o canal não é preenchido e o programa fica aguardando
	// o canal fica esperando (sleep)
	// Devemos tomar cuidado com o deadlock, go não identifica deadlock na compilação
	/*
		for{
			mensagem := <- canal
			fmt.Println(mensagem)
		}
	*/

	// Para resolver isso podemos usar um for
	/*
		for {
			mensagem, aberto := <-canal // Aguarda o canal ser preenchido

			// O segundo retorno do canal é um booleano que indica se o canal está aberto ou não

			if !aberto {
				break
			}

			fmt.Println(mensagem)
		}
	*/

	// Podemos usar o for range para ler o canal
	// Desse jeito não precisamos nos preocupar com o fechamento do canal
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}
