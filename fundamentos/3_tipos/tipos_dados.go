package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos numéricos inteiros
	// int8, int16, int32, int64

	fmt.Println("\nInteiros")

	var numero int8 = 100

	fmt.Println(numero)

	// tbm é possivel declarar somente int, nesse caso o compilador
	// vai definir o tamanho da variável de acordo com o sistema operacional

	var numero2 int = 1000000000000000000

	fmt.Println(numero2)

	// ou tbm com declaracao implicita
	// nesse caso sera definido como int64 sistema operacional de 64 bits
	numero3 := 1000000000000000000

	fmt.Println(numero3)

	// uint8, uint16, uint32, uint64
	// uint é um inteiro sem sinal, ou seja, não aceita valores negativos

	// var numero4 uint32 = -1000 // erro
	var numero4 uint32 = 1000

	fmt.Println(numero4)

	fmt.Println("\nAlias")
	// alias
	// INT32 = RUNE
	// Serve para representar um mapeamento da tabela Unicode (int32)
	var numero5 rune = 123456

	fmt.Println(numero5)

	// BYTE = UINT8
	// Serve para representar um mapeamento da tabela ASCII (uint8)
	var numero6 byte = 123

	fmt.Println(numero6)

	fmt.Println("\nNumeros reais")
	// Tipos numéricos reais
	// float32, float64

	var numeroReal1 float32 = 123.45

	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123564.45

	fmt.Println(numeroReal2)

	// declaracao implicita
	// nesse caso sera definido como float64 sistema operacional de 64 bits
	// mesmo caso do int
	numeroReal3 := 12345.67

	fmt.Println(numeroReal3)

	fmt.Println("\nString")

	// No go não existe o tipo char, somente string
	// sempre declarar com aspas duplas
	var str string = "Texto"

	fmt.Println(str)

	str2 := "Texto2"

	fmt.Println(str2)

	fmt.Println("\nCHAR")
	// tipo char
	// será impresso o valor da tabela ASCII
	// no fim do dia é um inteiro
	char := 'B'

	fmt.Println(char)

	fmt.Println("\nValor booleano")
	// valor booleano
	// true ou false
	var booleano bool = true

	fmt.Println(booleano)

	fmt.Println("\nValor zero")
	// valor zero
	// toda variável declarada tem um valor inicial
	// inteiros = 0
	// float = 0
	// string = ""
	// boolean = false
	// ponteiros = nil

	var texto string
	var numero7 int32
	var booleano1 bool

	// Funciona como valor default
	fmt.Println(texto, numero7, booleano1)

	fmt.Println("\nTipo erro")
	// Tipo erro
	// é um tipo de dado que representa um erro
	var erro error = errors.New("Erro interno")

	fmt.Println(erro)

	//por padrao o erro é nil
	// nil = null
	// nil é um valor default
	// nil é um valor que não aponta para nenhum endereço de memória
	var erro2 error
	fmt.Println(erro2)
}
