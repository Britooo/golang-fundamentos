package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Arrays
	// Assim como no Java, o array em Go é uma estrutura estática
	// Ou seja, o tamanho do array é fixo
	// O tamanho do array é definido na declaração
	// O array é uma estrutura homogênea (mesmo tipo de dados)
	// estrutura inflexivel
	var array1 [5]string

	array1[0] = "Posição 1"

	fmt.Println(array1)

	// Outra forma de declarar um array
	fmt.Println("\nOutra forma de declarar um array")
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	fmt.Println(array2)

	// O ... indica que o tamanho do array será definido pelo número de elementos
	// Na prática não é muito usado
	array3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array3)

	fmt.Println("\nSlice")

	// Slice
	// O slice é uma estrutura dinâmica
	// O slice é uma estrutura flexível
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	fmt.Println(slice1)

	// O slice é uma estrutura de dados que aponta para um array
	// Ele nao é um array

	// Provando que o slice aponta para um array
	fmt.Println("\nProvando que o slice aponta para um array")
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	// Para adicionar um elemento no slice, é preciso usar a função append
	slice1 = append(slice1, 20)

	fmt.Println(slice1)

	// se criarmos um novo slice, apontando para um array, e alterarmos o valor de um elemento
	// do array, o valor do elemento do slice também será alterado

	// Para pegar uma fatia de um array, podemos fazer assim:
	// O indice 1 é inclusivo
	// O indice 3 é exclusivo
	slice2 := array2[1:3]

	fmt.Println("\nSlice 2")
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"

	fmt.Println("\nSlice 2")
	fmt.Println(slice2)

	// desse jeito é um ponteiro

	fmt.Println("\nArrays Internos")
	// Arrays Internos
	// O slice aponta para um array interno
	// O array interno é criado automaticamente
	// Se nao passarmos o ultimo parametro, a capacidade será igual ao tamanho do slice
	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)
	fmt.Println("Tamanho ", len(slice3))   // tamanho do slice (length)
	fmt.Println("Capacidade", cap(slice3)) // capacidade do slice (capacity)

	// Se o slice atingir a capacidade máxima, o Go irá dobrar a capacidade do slice
	// comportamento semelhante ao ArrayList do Java

	slice3 = append(slice3, 3.5)
	slice3 = append(slice3, 5.5)

	fmt.Println(slice3)
	fmt.Println("Tamanho ", len(slice3))   // tamanho do slice (length)
	fmt.Println("Capacidade", cap(slice3)) // capacidade do slice (capacity)

	//No fim do dia, quando declaramos desse jeito:
	var slice4 []int
	// Estamos criando um slice, que aponta para um array interno

	slice4 = append(slice4, 18)

	fmt.Println(slice4)
}
