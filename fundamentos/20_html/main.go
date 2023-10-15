package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// O pacote "html/template" é usado para renderizar templates HTML.
var templates *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func main() {

	usuario := Usuario{
		Nome:  "Ludmilo",
		Email: "ludmilo@cnpjoto",
	}

	// O pacote "html/template" é usado para renderizar templates HTML.
	// O método ParseGlob é usado para carregar todos os templates de um diretório.
	// O método Must é usado para tratar erros de parsing.
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// O método ExecuteTemplate é usado para executar um template específico.
		// O primeiro argumento é o ResponseWriter, o segundo é o nome do template e o terceiro é o dado que será passado para o template.
		templates.ExecuteTemplate(w, "index.html", usuario)
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
