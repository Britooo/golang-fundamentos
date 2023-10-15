package main

import (
	"fmt"
	"time"
)

// Padrão Generator:
// O padrão Generator é um padrão de design utilizado em Go onde uma função cria (gera) um canal e
// inicia uma goroutine que envia valores para esse canal. A função então retorna o canal para o chamador,
// permitindo ao chamador receber valores desse canal de maneira fácil e encapsulada.

// escrever é uma implementação do padrão Generator.
// Ela cria e retorna um canal onde strings são enviadas repetidamente.
// A string enviada é a concatenação da mensagem "Valor recebido" com o `texto` passado como parâmetro.
func escrever(texto string) <-chan string {
	canal := make(chan string) // Criando um canal de strings.

	// Esta goroutine envia repetidamente mensagens para o canal.
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto) // Enviando a mensagem para o canal.
			time.Sleep(time.Millisecond * 500)               // Pausando por 500 milissegundos (meio segundo) antes de enviar novamente.
		}
	}()

	return canal // Retornando o canal.
}

func main() {

	fmt.Println("Padrão Generator")

	// Chama a função escrever, um exemplo do padrão Generator, e obtém o canal retornado por ela.
	canalMain := escrever("Generator 1")

	// Lê e imprime 10 mensagens do canal.
	for i := 0; i < 10; i++ {
		fmt.Println(<-canalMain)
	}
}
