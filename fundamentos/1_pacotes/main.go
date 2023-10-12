package main

import (
	"modulo/auxiliar"
)

// Sobre modulos:
// https://golang.org/doc/tutorial/create-module
// https://blog.golang.org/using-go-modules
// Podemos usar o comando go mod init para criar um arquivo go.mod
// que irá conter o nome do módulo e a versão do Go que estamos usando.
// O comando go mod tidy irá adicionar os pacotes que estamos usando
// no arquivo go.mod e remover os que não estamos usando.
// Isso é útil para manter o arquivo go.mod atualizado.
func main() {
	auxiliar.Escrever("Olá Mundo!")
}
