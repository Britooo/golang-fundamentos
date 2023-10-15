package endereco

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoEndereco verifica o tipo de um logradouro informado
func TipoEndereco(texto string) string {

	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoNormalizado := strings.ToLower(texto)

	primeiraPalavra := strings.Split(enderecoNormalizado, " ")[0]

	caser := cases.Title(language.BrazilianPortuguese)

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			return caser.String(tipo)
		}
	}

	return "Tipo inválido"
}
