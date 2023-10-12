package main

import "fmt"

func main() {
	// Duas maneiras de declarar variáveis
	// A primeira maneira, declaração explícita
	var nome1 string = "Diego"

	// A segunda maneira é mais utilizada, declaração implícita
	nome2 := "João"

	fmt.Println(nome1, nome2)

	// Declaração de variáveis em bloco explícito
	var (
		nome3 string = "José"
		nome4 string = "Maria"
	)

	fmt.Println(nome3, nome4)

	// Declaração de variáveis em bloco implícito
	nome5, nome6 := "Pedro", "Lucas"

	fmt.Println(nome5, nome6)

	//Constantes
	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	// Invertendo valores de variáveis
	nome1, nome2 = nome2, nome1

	fmt.Println(nome1, nome2)
}
