package main

import "fmt"

func main() {
	// Operadores aritméticos
	// + - / * %

	fmt.Println("Operadores aritméticos")

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Variáveis de tipo diferente não podem ser operadas
	// Exemplo:

	// var numero1 int16 = 10
	// var numero2 int32 = 25
	// soma2 := numero1 + numero2

	fmt.Println("Operadores de atribuição")

	// Operadores de atribuição
	// := = += -= *= /= %=
	var numero int16 = 10
	numero += 10 // numero = numero + 10

	fmt.Println(numero)

	fmt.Println("Operadores relacionais")

	// Operadores relacionais
	// == != > < >= <=

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)

	idade := 18
	maiorDeIdade := idade >= 18

	fmt.Println(maiorDeIdade)

	fmt.Println("Operadores lógicos")

	// Operadores lógicos
	// && (AND)
	// || (OR)
	// ! (NOT)

	temHabilitacao := true
	maiorDeIdade2 := idade >= 18

	fmt.Println("Operador AND")
	fmt.Println(temHabilitacao && maiorDeIdade2)

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Operadores unários
	// ++ --
	fmt.Println("Operadores unários")

	numero3 := 10

	numero3++

	fmt.Println(numero3)

	// Não existe pre-decremento e pre-incremento em go
	// --numero3 ++numero3
	numero3--

	fmt.Println(numero3)

	// Não existe operador ternário em go
	// Exemplo:
	// texto := numero3 > 5 ? "Maior que 5" : "Menor que 5"

	var texto string

	if numero3 > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

	// No Go, o operador ternário tem uma premissa
	// Só tem uma maneira fazer as coisas
	// exceto em declaracao de variáveis
}
