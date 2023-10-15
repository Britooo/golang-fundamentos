package main

import (
	"fmt"
	"math/rand"
	"time"
)

// escrever é uma implementação do padrão Generator.
// Ela cria e retorna um canal onde strings são enviadas repetidamente.
func escrever(texto string) <-chan string {
	canal := make(chan string) // Criando um canal de strings.

	// Esta goroutine envia repetidamente mensagens para o canal.
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			// Introduz uma pausa aleatória antes de enviar a próxima mensagem.
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal // Retornando o canal.
}

// Padrão Multiplexador:
// O padrão Multiplexador é um padrão de design utilizado em Go para combinar vários canais em um único canal.
// Isso é útil quando você quer ouvir múltiplos canais e processar as mensagens recebidas de qualquer um deles de maneira uniforme.

// Multiplexar combina as mensagens de dois canais de entrada em um único canal de saída.
func Multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string) // Canal que irá combinar as mensagens de ambos os canais de entrada.

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalDeSaida <- mensagem // Envia a mensagem do canal1 para o canal de saída.
			case mensagem := <-canal2:
				canalDeSaida <- mensagem // Envia a mensagem do canal2 para o canal de saída.
			}
		}
	}()

	return canalDeSaida
}

func main() {
	fmt.Println("Multiplexador")

	// Utilizando o Multiplexador para combinar mensagens de dois canais gerados pela função escrever.
	canalMain := Multiplexar(escrever("Mensagem 1"), escrever("Mensagem 2"))

	// Lê e imprime 10 mensagens do canal multiplexado.
	for i := 0; i < 10; i++ {
		fmt.Println("MSG: ", <-canalMain)
	}
}
