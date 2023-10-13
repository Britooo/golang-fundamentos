package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {

	// For infinito
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Go routines")

	// Concorrencia != Paralelismo

	// Concorrencia: Executar não necessiarimente ao mesmo tempo, revezando entre as tarefas
	// Paralelismo: Executar várias tarefas ao mesmo tempo em vários processadores

	//1. TESTE sem concorrencia, descomentar as linhas abaixo
	// escrever("Olá Mundo")
	// escrever("Programando em Go")

	// Nesse caso, nunca será executado a segunda função, pois a primeira nunca termina

	// GO ROUTINES
	// É uma forma de executar funções de forma concorrente
	// Para executar uma função de forma concorrente, basta colocar a palavra go antes da chamada da função

	// 2. TESTE com concorrencia, descomentar as linhas abaixo
	// Independente da primeira chamada terminar ou não, a segunda será executada
	go escrever("Olá Mundo")
	escrever("Programando em Go")

	/*
		// 3. TESTE com concorrencia, descomentar as linhas abaixo
		go escrever("Olá Mundo")
		go escrever("Programando em Go")

		Cuidado! nesse caso, o programa termina antes de executar as funções
	*/
}
