package main

import "fmt"

func main() {

	fmt.Println("Estruturas de controle")

	// Estruturas de controle

	numero := 10

	// If - else if - else
	if numero < 15 {
		fmt.Println("Menor que 15")
	} else if numero == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Maior que 15")
	}

	// Não funciona sem chaves
	// if numero > 5
	// 	fmt.Println("Maior que 5")

	// Não requer parênteses
	// if (numero > 5) {
	// 	fmt.Println("Maior que 5")
	// }

	fmt.Println("If init")

	// If init
	// if init; condition {}
	// A variavel declarada no init só existe dentro do escopo do if
	// Útil para inicializar variáveis perante uma condicional
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")

		fmt.Println(outroNumero)
	} else {
		outroNumero := 0
		fmt.Println(outroNumero)
		fmt.Println("Numero é menor que zero")
	}

	// Não funciona
	// fmt.Println(outroNumero)
}
