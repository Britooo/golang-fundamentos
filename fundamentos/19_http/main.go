package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Protocolo HTTP")

	// O protocolo HTTP (HyperText Transfer Protocol) é o protocolo utilizado para transferência de dados na web.
	// Ele opera no nível de aplicação e é baseado em um modelo de requisição-resposta.

	// Requisição HTTP:
	// - Método: Indica a operação desejada (GET, POST, PUT, DELETE, etc.).
	// - URL: Endereço do recurso que está sendo solicitado.
	// - Headers: Metadados da requisição, como tipo de conteúdo, autenticação, entre outros.
	// - Body: Corpo da requisição, usado para enviar dados ao servidor (especialmente em POSTs e PUTs).

	// Resposta HTTP:
	// - Status: Código numérico que indica o resultado da requisição (200 para sucesso, 404 para não encontrado, etc.).
	// - Headers: Metadados da resposta, como tipo de conteúdo, tamanho do conteúdo, entre outros.
	// - Body: Corpo da resposta, contém os dados retornados pelo servidor.

	// Métodos HTTP comuns:
	// - GET: Recuperar informações sobre um recurso.
	// - POST: Enviar dados para serem processados por um recurso.
	// - PUT: Atualizar um recurso existente ou criar um novo.
	// - DELETE: Remover um recurso.

	// Em navegadores web, o método GET é o mais comum, utilizado para solicitar páginas web.

	// A função http.HandleFunc é usada para registrar uma função de manipulação (handler) para um path específico.
	// Neste caso, estamos registrando uma função anônima para o path "/".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo!"))
	})

	// Aqui, registramos a função "teste" para o path "/teste".
	http.HandleFunc("/teste", teste)

	// A função http.ListenAndServe inicia um servidor HTTP que escuta em um determinado endereço e porta.
	// No caso ":8080", o servidor escuta em todas as interfaces de rede na porta 8080.
	// O segundo argumento é o manipulador; se for nil, ele usa DefaultServeMux.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Função "teste" que é chamada quando um cliente faz uma requisição para "/teste".
func teste(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste"))
}
