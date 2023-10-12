package main

// Retorno nomeado
// func nomeFuncao(parametros) (retorno tipo) {}
// func nomeFuncao(parametros) (retorno1 tipo1, retorno2 tipo2) {}
// Serve para quando temos muitos retornos
// Declaramos quais variaveis serão retornadas, sendo assim não precisamos
// declarar o retorno + variavel
func operacoesAritmeticas(n1 int, n2 int) (soma int, subtracao int) {

	// Na declaração não precisamos colocar := pois já declaramos o tipo
	// das variaveis no retorno da função

	soma = n1 + n2
	subtracao = n1 - n2

	// Não precisamos colocar o return pois já declaramos as variaveis
	// de retorno na declaração da função
	// return soma, subtracao -> seria redundante
	return
}

func main() {

	soma, subtracao := operacoesAritmeticas(10, 20)

	println(soma, subtracao)
}
