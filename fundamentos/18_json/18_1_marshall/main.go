// Este é um programa que demonstra a manipulação de estruturas JSON em Go.

package main

import (
	"bytes"         // Pacote utilizado para operações de buffer.
	"encoding/json" // Pacote utilizado para a codificação e decodificação JSON.
	"fmt"
	"log" // Pacotes para saída padrão e log.
)

// Definindo uma struct chamada Doguinho com tags JSON.
// As tags são usadas para mapear os campos da struct para as respectivas chaves JSON quando é codificado.
type Doguinho struct {
	Nome  string `json:"nome"`  // O campo "Nome" será representado como "nome" em JSON.
	Raca  string `json:"raca"`  // O campo "Raca" será representado como "raca" em JSON.
	Idade uint   `json:"idade"` // O campo "Idade" será representado como "idade" em JSON.
}

func main() {

	// Mensagem inicial para indicar a operação principal do programa.
	fmt.Println("Manipulando JSON")

	// Criando uma instância da struct Doguinho.
	dog1 := Doguinho{"Caramelo, o doguinho amarelo", "Vira-lata", 3}

	// Convertendo a struct dog1 em uma representação JSON.
	// A função json.Marshal retorna dois valores: o JSON como um slice de bytes e um erro (caso ocorra).
	doguinhoJSON, error := json.Marshal(dog1)

	// Verificando se ocorreu algum erro durante a codificação.
	if error != nil {
		log.Fatal(error) // Se houver um erro, ele será registrado e o programa será encerrado.
	}

	// Convertendo o slice de bytes em uma string para exibição mais legível.
	// O buffer permite trabalhar com os bytes como se fossem uma string.
	result := bytes.NewBuffer(doguinhoJSON)

	// Imprimindo a representação JSON de dog1.
	fmt.Println(result)

	// Exemplo mostrando que é possível converter um map para JSON.
	// Criando um map representando outro "doguinho".
	dog2 := map[string]string{
		"nome": "Mil grau",
		"raca": "Vira-lata",
	}

	// Convertendo o map dog2 em uma representação JSON.
	doguinho2JSON, error := json.Marshal(dog2)

	// Verificando se ocorreu algum erro durante a codificação.
	if error != nil {
		log.Fatal(error) // Se houver um erro, ele será registrado e o programa será encerrado.
	}

	// Convertendo o slice de bytes em uma string para exibição mais legível.
	resultado2 := bytes.NewBuffer(doguinho2JSON)

	// Imprimindo a representação JSON de dog2.
	fmt.Println(resultado2)
}
