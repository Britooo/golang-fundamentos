package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Estruturas de repetição")

	//Go lang somente tem o for
	// Não tem while, do while, foreach
	// Não tem parênteses

	numero := 0

	// Algo próximo ao while
	// for condition {}
	for numero < 10 {
		numero++
		time.Sleep(time.Second)
		fmt.Printf("Incrementando o numero: %d\n", numero)
	}

	fmt.Println("Outra maneira de fazer o for, parecido com if init")

	// Abordagem mais próxima do fori tradicional
	// for init; condition; pos {}
	for j := 0; j < 10; j += 2 {
		time.Sleep(time.Second)
		fmt.Printf("Incrementando o numero: %d\n", j)
	}

	fmt.Println("For infinito")

	// For infinito
	// for {}
	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }

	fmt.Println("For range (clausula range)")

	nomes := []string{"João", "Davi", "Lucas"}

	// Algo semelhante ao foreach
	// for index, value := range array {}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	/*
		for nome := range nomes {
			fmt.Println(nome)

			Até daria certo mas a variavel nome seria o indice :/
		}
	*/

	// Se não precisar do indice, pode fazer assim
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("Iterando sobre string")

	var nomeUsuario string = "Diego"

	// Iterando sobre string

	/*
			for indice, letra := range "Texto" {
			fmt.Println(indice, string(letra))
		}
	*/

	fmt.Println("Iterando sobre string ASCII")

	// Isso imprimiria o valor em bytes (ASCII)
	for indice, letra := range nomeUsuario {
		fmt.Println(indice, letra)
	}

	fmt.Println("Iterando sobre string Unicode")
	// para pegar o valor em texto, é preciso converter para string
	for indice, letra := range nomeUsuario {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Diego",
		"sobrenome": "Brito",
	}

	fmt.Println("Iterando sobre map")

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
