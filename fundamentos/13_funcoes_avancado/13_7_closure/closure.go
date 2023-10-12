package main

import "fmt"

// Em resumo, uma closure é uma função que referencia variáveis externas
// A função pode acessar e atribuir as variáveis externas
// Em outras palavras, a função "lembra" do ambiente onde ela foi criada
func exemploClosure() func() {
	texto := "Dentro da função exemploClosure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	fmt.Println("Closure")

	// Mesmo que tenhamos uma variável com o mesmo nome
	// a função exemploClosure() irá acessar a variável externa
	// e não a variável interna

	texto := "Dentro da função main"

	fmt.Println(texto)

	funcaoClosure := exemploClosure()

	funcaoClosure()
}
