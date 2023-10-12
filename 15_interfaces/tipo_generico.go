package main

import "fmt"

func generica(i interface{}) {
	fmt.Println(i)
}

func main() {

	fmt.Println("Tipo genérico - Interface vazia - CUIDADO!!!")

	// Uma interface vazia pode ser usada como um tipo genérico
	// Isso é útil quando não sabemos o tipo de dado que será recebido
	// Porém, é preciso tomar cuidado ao usar uma interface vazia
	// Pois não há garantia de que o tipo recebido será o esperado

	generica("String")
	generica(1)
	generica(true)

	// O comando fmt.Println() também usa uma interface vazia,
	// Na verdade, recebe n parâmetros do tipo interface{},
	// e imprime o valor de cada um deles,
	// é considerada uma função variática

	// Abaixo um péssimo exemplo de uso de interface vazia
	// Seria algo semelhante a um Object em Java
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         1,
	}

	fmt.Println(mapa)

	// Vale lembrar que Go é uma linguagem estática
	// fortemente tipada, não é uma linguagem dinâmica
	// como Python, PHP, etc
	// Por isso, não é muito comum o uso de interface vazia
	// em Go, pois perdemos a segurança de tipos
	// TOME CUIDADO AO USAR INTERFACE VAZIA!!!
}
