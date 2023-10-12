package main

import "fmt"

// Funções recursivas
// São funções que chamam a si mesmas
// Não recomendadas para casos com muitas repetições
// Exemplo de função recursiva:
func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	// A função fibonacci chama a si mesma
	// fibonacci(8) -> fibonacci(6) + fibonacci(7)
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	fmt.Println("Funções recursivas")

	// Uma função recursiva é uma função que chama a si mesma
	// Exemplo:

	// Clássico exercício de Fibonacci
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55...

	resultado := fibonacci(8)

	fmt.Println(resultado)
}
