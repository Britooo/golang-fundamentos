package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type estudante struct {
	// Seria redundante declarar novamente os atributos de pessoa
	// nome      string
	// sobrenome string
	// idade     int
	// Então podemos fazer uma relação de "herança" entre structs
	// Unico caso que nao precisamos declarar o tipo da struct é quando ela é um ponteiro
	pessoa
	curso     string
	faculdade string
}

func main() {
	// Em Go não existe herança de classes, mas é possível criar uma relação de herança entre structs

	fmt.Println("Herança - SQN")

	p1 := pessoa{"Diego", "Brito", 30}
	fmt.Println(p1)

	estudante1 := estudante{p1, "Administração", "USP"}

	fmt.Println(estudante1)

	// Acessando atributos de pessoa
	// poderia ser estudante1.pessoa.nome mas Go permite omitir o nome da struct
	// por conta da declaração
	fmt.Println(estudante1.nome)
}
