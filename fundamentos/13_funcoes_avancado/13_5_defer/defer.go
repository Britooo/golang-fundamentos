package main

import "fmt"

func funcao1() {
	println("Executando a função 1")
}

func funcao2() {
	println("Executando a função 2")
}

func verificarSituacaoAluno(nota1 float64, nota2 float64) bool {

	// A cláusula defer adia a execução de uma função
	// até que a função externa retorne
	// Neste caso, a função fmt.Println("Média calculada. Resultado será retornado!")
	// será executada após a função verificarSituacaoAluno retornar
	// Evitando assim, a repetição de código
	defer fmt.Println("Média calculada. Resultado será retornado!")

	fmt.Println("Entrando na função para verificar a situação do aluno")

	media := (nota1 + nota2) / 2

	return media >= 6
}

func main() {

	fmt.Println("Cláusula Defer")

	// A cláusula defer adia a execução de uma função
	// A função adiada é executada mesmo que ocorra algum erro
	// Exemplo:

	//defer funcao1()
	//funcao2()

	// A cláusula defer é muito utilizada para fechar conexões com banco de dados
	// ou fechar arquivos abertos :)

	resultado := verificarSituacaoAluno(7, 8)

	if resultado {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
