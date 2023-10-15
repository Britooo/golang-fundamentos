// O pacote "forma_test" é utilizado para testar funcionalidades relacionadas às formas geométricas.
package forma_test

import (
	"math"
	"testing"

	// Importando os pacotes necessários. "math" contém constantes e funções matemáticas, e "testing" é utilizado para criar testes unitários em Go.

	// O "." antes de "teste-avancado/forma" permite acessar funções e tipos do pacote "forma" diretamente, sem o uso de prefixo.
	. "teste-avancado/forma"
)

// Definindo uma struct para organizar cenários de teste.
// FormaGeometrica é provavelmente uma interface que tem um método "Area".
type cenariosDeTeste struct {
	FormaGeometrica // Incorporando a interface (ou tipo) FormaGeometrica.
	areaEsperada    float64
}

// A função "TestArea" testa a funcionalidade de cálculo de área para diferentes formas geométricas.
func TestArea(t *testing.T) {

	// O conceito de "subtestes" em Go permite organizar testes para uma mesma funcionalidade por diferentes casos ou cenários.
	// Aqui, estamos usando "t.Run" para criar um subteste para a forma "Retângulo".
	t.Run("Retângulo", func(t *testing.T) {

		// Definindo cenários de teste específicos para retângulos.
		cenarios := []cenariosDeTeste{
			{Retangulo{10, 12}, 120.0},
			{Retangulo{10, 6}, 60.0},
			{Retangulo{12, 6}, 72.0},
			{Retangulo{12, 12}, 144.0},
		}

		// Iterando sobre os cenários de teste.
		for _, cenario := range cenarios {
			// Calculando a área usando o método "Area" da interface "FormaGeometrica".
			areaRecebida := cenario.Area()

			// Verificando se a área calculada é igual à esperada.
			if areaRecebida != cenario.areaEsperada {
				// "t.Fatalf" registra um erro e interrompe a execução do teste atual.
				t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, cenario.areaEsperada)
			}
		}
	})

	// Criando um subteste para a forma "Círculo".
	t.Run("Círculo", func(t *testing.T) {

		// Definindo cenários de teste específicos para círculos.
		cenariosDeTeste := []cenariosDeTeste{
			{Circulo{10}, math.Pi * 100},
			{Circulo{5}, math.Pi * 25},
			{Circulo{2}, math.Pi * 4},
			{Circulo{3}, math.Pi * 9},
		}

		// O restante do código segue a mesma lógica do subteste "Retângulo":
		// - Calcula a área para cada círculo nos cenários de teste
		// - Compara com a área esperada
		// - Registra um erro se as áreas não coincidirem
		for _, cenario := range cenariosDeTeste {

			areaRecebida := cenario.Area()

			if areaRecebida != cenario.areaEsperada {

				t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, cenario.areaEsperada)
			}
		}
	})
}
