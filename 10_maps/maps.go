package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Maps
	// Maps sao estruturas de dados que armazenam valores associados a chaves
	// Maps sao tambem conhecidos como dicionarios ou hashes em outras linguagens
	// Dentro do colchetes, temos o tipo da chave e fora o tipo do valor

	user := map[string]string{
		"name":     "John",
		"lastName": "Doe",
	}

	fmt.Println(user)

	// Acessando um valor
	fmt.Println("Acessando um valor")
	// Nao funciona com . (ponto)
	// fmt.Println(user.name)
	fmt.Println(user["name"])

	fmt.Println("Adicionando um map dentro de outro")
	// Um map pode ter outros maps dentro dele
	user2 := map[string]map[string]string{
		"name": {
			"first":    "John",
			"lastName": "Doe",
		},
		"course": {
			"name": "Golang",
		}, // Sempre deixe uma virgula no ultimo elemento
	}

	fmt.Println(user2["name"]["first"])
	fmt.Println(user2)

	fmt.Println("removendo um elemento")
	// Removendo um elemento
	delete(user2, "course")

	fmt.Println(user2)

	fmt.Println("Adicionando um elemento")
	// Adicionando um elemento
	user2["age"] = map[string]string{
		"age": "30",
	}

	fmt.Println(user2)
}
