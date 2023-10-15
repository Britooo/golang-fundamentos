package main

import (
	"fmt"
	"introducao-testes/endereco"
)

func main() {
	tipoEndereco := endereco.TipoEndereco("Avenida Paulista")

	fmt.Printf("Resultado: %s", tipoEndereco)
}
