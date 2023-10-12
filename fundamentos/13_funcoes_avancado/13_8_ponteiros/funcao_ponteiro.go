package main

import "fmt"

func inverterSinalNumero(numero int) int {
	return numero * -1
}

func inverterSinalNumeroPonteiro(numero *int) {
	*numero = *numero * -1 // Desreferenciando o ponteiro para acessar o valor e não o endereço de memória (*)
}

func main() {

	fmt.Println("Funções com ponteiros")

	numero1 := 20

	// Nesse caso estamos enviando o valor da variável numero1
	// Ou seja, estamos enviando uma cópia da variável
	// A função inverterSinalNumero() irá trabalhar com uma cópia da variável
	// Portanto, a variável numero1 não será alterada
	numeroInvertido := inverterSinalNumero(numero1)

	fmt.Printf("Número: %d, Invertido: %d\n", numero1, numeroInvertido)

	numero2 := 40

	// Nesse caso estamos enviando o endereço de memória da variável numero2
	// Ou seja, estamos enviando o ponteiro da variável
	inverterSinalNumeroPonteiro(&numero2)

	fmt.Printf("Número: %d, Invertido: %d\n", numero2, numero2)
}
