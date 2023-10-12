package main

import "fmt"

func main() {
	resultado := soma(1, 2)

	fmt.Println(resultado)
	escreve("Olá Mundo!")

	// Funções anônimas
	// São funções sem nome, que podem ser atribuídas a variáveis
	var funcaoAnonima = func(texto string) string {
		return texto
	}

	resultado2 := funcaoAnonima("Função anônima!")

	fmt.Println(resultado2)

	fmt.Println("\nFunções com mais de um retorno")
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)

	fmt.Println(resultadoSoma, resultadoSubtracao)

	fmt.Println("\nFunções com mais de um retorno - parte 2")
	fmt.Println("\nIgnorando um dos retornos")
	// Ignorando um dos retornos
	// Se nao quiser usar um dos retornos, basta usar o underline
	resultadoSoma2, _ := calculosMatematicos(10, 20)

	fmt.Println(resultadoSoma2)
}

// Em go, funcoes podem retornar mais de um valor
// Exemplo:
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// Em go, o tipo de retorno da função é declarado após os parênteses
// Exemplo:
func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Se nao tiver retorno, nao precisa declarar o tipo de retorno
func escreve(texto string) {
	fmt.Println(texto)
}
