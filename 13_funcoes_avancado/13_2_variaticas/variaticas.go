package main

import "fmt"

// Funções variáticas
// São funções que recebem uma quantidade variável de parâmetros
func soma(numeros ...int) (resultado int) {

	// ...int -> significa que a função pode receber uma quantidade variável
	// de parâmetros do tipo int (slice de int)

	for _, numero := range numeros {
		resultado += numero
	}

	return
}

// Podemos combinar parametros variáticos com parametros normais
// Porém, os parametros variáticos devem ser os últimos
// Somente um por função
// Exemplo:
func escrever(texto string, numeros ...int) {

	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	fmt.Println("Funções variáticas")

	resultado := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(resultado)

	fmt.Println("Funções variáticas sem parametros")
	resultado = soma()

	fmt.Println(resultado)

	fmt.Println("Funções variáticas com parametros normais")
	escrever("Olá mundo", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
