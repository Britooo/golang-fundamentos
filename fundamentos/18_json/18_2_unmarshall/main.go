// O pacote main é a entrada padrão para um aplicativo Go.
package main

// Importando os pacotes necessários.
import (
	"encoding/json" // Pacote padrão do Go para manipulação de JSON.
	"fmt"           // Pacote padrão do Go para operações de entrada/saída formatada.
	"log"           // Pacote padrão do Go para logging.
)

// Definição da struct Catioro.
type Catioro struct {
	Nome  string `json:"nome"`  // Este campo será mapeado para "nome" em JSON.
	Raca  string `json:"-"`     // O sinal "-" indica que este campo deve ser ignorado ao realizar marshal/unmarshal.
	Idade int    `json:"idade"` // Este campo será mapeado para "idade" em JSON.
}

func main() {

	// Uma string contendo um JSON representando um Catioro.
	catioro1JSON := `{"nome":"Calabreso","raca":"Vira-lata","idade":3}`

	// Declaração de uma variável do tipo Catioro.
	var catioro1 Catioro

	// Tentativa de unmarshal (converter) a string JSON para a struct Catioro.
	if error := json.Unmarshal([]byte(catioro1JSON), &catioro1); error != nil {
		// Se houver um erro durante a conversão, o programa logará o erro e terminará.
		log.Fatal(error)
	}

	// Impressão dos dados do Catioro 1.
	fmt.Println("\nCatioro 1")
	fmt.Println(catioro1)

	// Uma string contendo um JSON representando um Catioro.
	catioro2JSON := `{"nome":"Domitilo","raca":"Poodle"}`

	// Declaração de uma variável do tipo map para armazenar os dados do Catioro.
	var catioro2 map[string]string

	// Tentativa de unmarshal (converter) a string JSON para o map.
	if error := json.Unmarshal([]byte(catioro2JSON), &catioro2); error != nil {
		// Se houver um erro durante a conversão, o programa logará o erro e terminará.
		log.Fatal(error)
	}

	// Impressão dos dados do Catioro 2.
	fmt.Println("\nCatioro 2")
	fmt.Println(catioro2)
}
