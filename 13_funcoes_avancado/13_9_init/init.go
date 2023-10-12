package main

import "fmt"

var numero int

// A função init() é executada antes da função main()
// Podemos ter apenas uma função init() por arquivo, não por pacote
// A função init() é executada apenas uma vez
// Utilizada para inicializar variáveis, conexões com banco de dados, etc
func init() {
	fmt.Println("Inicializando...")
	numero = 10
}

func main() {
	fmt.Println("Main...")
	fmt.Println(numero)
}
