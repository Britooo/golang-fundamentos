package main

import "fmt"

// O mais próximo de classes que temos em Go são as structs
// Structs são tipos de dados que permitem armazenar diferentes tipos de dados
// Uma coleção de campos
// Exemplo:

type usuario struct {
	nome  string
	idade uint8
	// podemos ter structs dentro de structs
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	var user usuario

	//Valor zero de um struct: atributos inicilizados com seus valores zero
	fmt.Println(user)

	user.nome = "Diego"
	user.idade = 30

	fmt.Println(user)

	//Exemplo com endereco

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	// Outra forma de declarar um struct
	user2 := usuario{"Diego", 30, enderecoExemplo}

	fmt.Println(user2)

	// Outra forma de declarar um struct
	// Em casos de nao ter todos os atributos, os atributos nao declarados recebem o valor zero
	user3 := usuario{nome: "Diego"}

	fmt.Println(user3)
}
