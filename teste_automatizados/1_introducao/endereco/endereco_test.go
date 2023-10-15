// Em Go, ao criar um pacote de testes, você adiciona o sufixo "_test" ao nome do pacote.
// Isso ajuda a separar o código real dos testes, garantindo que os testes não sejam incluídos em builds de produção.
package endereco_test

import (
	"testing"
	// "testing" é o pacote padrão em Go para criar e executar testes unitários.

	// A notação "." antes do import é um recurso de Go chamado "dot import".
	// Ele permite que você acesse funções e variáveis do pacote importado diretamente, sem prefixo.
	// No entanto, é geralmente usado com cautela, pois pode tornar o código menos claro.
	. "introducao-testes/endereco"
)

// Definindo uma struct para representar cenários de teste. Isso é útil para testar várias entradas
// e saídas esperadas em uma única função de teste.
type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Funções de teste em Go devem começar com a palavra "Test" seguido do nome da função que você deseja testar.
func TestTipoEndereco(t *testing.T) {

	// t.Parallel() pode ser usado para executar testes em paralelo, aproveitando a concorrência e acelerando a execução.
	// No entanto, tenha cuidado ao usar isso com testes que acessam recursos compartilhados.
	// t.Parallel()

	// Aqui, estamos definindo diversos cenários para testar a função "TipoEndereco".
	cenariosDeTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	// Iteramos sobre cada cenário de teste.
	for _, cenario := range cenariosDeTeste {
		// Para cada cenário, chamamos a função que queremos testar e comparamos o resultado ao esperado.
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			// t.Errorf é usado para registrar um erro no teste. A execução do teste continuará mesmo depois de registrar um erro.
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido, cenario.retornoEsperado)
		}
	}
}

// Função de teste simples que demonstra um teste que não falha.
func TestAleatorio(t *testing.T) {
	// Para rodar em parelelo.
	// t.Parallel()

	if false {
		t.Error("Teste falhou")
	}
}
