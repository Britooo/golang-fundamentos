package main

import "fmt"

// fibonacci retorna o enésimo número da sequência Fibonacci.
func fibonacci(n int) int {
	if n <= 1 {
		return n // Se n for 0 ou 1, retorna o próprio n.
	}
	// Retorna a soma dos dois números anteriores na sequência Fibonacci.
	return fibonacci(n-1) + fibonacci(n-2)
}

// worker é uma função que "ouve" as tarefas do canal tarefas,
// processa a tarefa e envia o resultado para o canal resultados.
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	fmt.Println("Worker Pools")

	// Criando dois canais: um para tarefas e outro para resultados. Ambos com capacidade para 45 itens.
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Iniciando 4 workers (goroutines) em paralelo.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Enviando números de 0 a 44 para o canal de tarefas.
	for i := 0; i < cap(tarefas); i++ {
		tarefas <- i
	}

	// Fechando o canal de tarefas para indicar que não serão adicionadas mais tarefas.
	close(tarefas)

	// Lendo e imprimindo os resultados do canal de resultados.
	for i := 0; i < cap(resultados); i++ {
		resultado := <-resultados
		fmt.Println("Resultado: ", resultado)
	}

	// Fechando o canal de resultados.
	close(resultados)
}
