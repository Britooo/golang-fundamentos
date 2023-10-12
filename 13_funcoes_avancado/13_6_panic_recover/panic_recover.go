package main

import "fmt"

func recuperarExecucao() {

	// A função recover recupera a execução do programa em caso de erro
	// Exemplo:

	// Caso não haja erro, a função recover retorna nil
	// Caso haja erro, a função recover retorna o erro
	if recuperacao := recover(); recuperacao != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}

	// Desta forma, o programa não é interrompido, ou seja, não para de executar
}

func verificarSituacaoAluno(nota1, nota2 float64) bool {

	// Para ver o funcionamento do panic e recover, comente a linha abaixo
	defer recuperarExecucao()

	media := (nota1 + nota2) / 2

	if media >= 6 {
		return true
	} else if media >= 0 && media < 6 {
		return false
	}

	// A função panic interrompe a execução do programa
	// Literalmente, o programa entra em pânico
	// Se não tratarmos o erro, o programa será interrompido
	// A cláusula recover recupera a execução do programa
	panic("A média é negativa!")
}

func main() {

	fmt.Println("Panic e Recover")

	// A função panic interrompe a execução do programa
	// A função recover recupera a execução do programa
	// Exemplo:

	resultado := verificarSituacaoAluno(10, 10)

	if resultado {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
