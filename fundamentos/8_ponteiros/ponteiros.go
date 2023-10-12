package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var numero1 int = 10
	var numero2 int = numero1

	numero1++

	fmt.Println("Qual váriavel mudou:")
	fmt.Printf("Número 1: %d\nNúmero 2: %d\r\n", numero1, numero2)

	// Por padrao Go sempre criar uma copia da variavel
	// Para alterar o valor da variavel original, é preciso usar ponteiros
	// Ponteiro é uma referencia de memoria
	var numero3 int = 100

	// *int é o tipo do ponteiro
	var ponteiro *int

	numero3 = 1000

	ponteiro = &numero3

	fmt.Println("\nQual váriavel mudou:")
	fmt.Printf("Número 3: %d\nPonteiro: %d\r\n", numero3, ponteiro)

	// Para exibir o valor do ponteiro, é preciso usar o *
	fmt.Println("\nQual váriavel mudou:")
	fmt.Printf("Número 3: %d\nPonteiro: %d\r\n", numero3, *ponteiro)

	// Esse processo é chamado de desreferenciação
	// em inglês é chamado de dereference

	// Por que usar ponteiros?

	numero3 = 42

	// Endereço de memoria nao muda, pois é uma referencia
	fmt.Println("\nQual váriavel mudou:")
	fmt.Printf("Número 3: %d\nPonteiro: %d\r\n", numero3, ponteiro)

	fmt.Println("\nQual váriavel mudou:")
	fmt.Printf("Número 3: %d\nPonteiro: %d\r\n", numero3, *ponteiro)
}
