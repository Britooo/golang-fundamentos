package forma

import "math"

// FormaGeomtrica é uma interface que define o comportamento de uma forma geométrica.
type FormaGeometrica interface {
	Area() float64
}

// Area calcula a área de uma forma geométrica.
func Area(forma FormaGeometrica) float64 {
	return forma.Area()
}

// Area calcula a área de um retângulo.
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Area calcula a área de um círculo.
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2) //(c.raio * c.raio)
}

// Retangulo representa um retângulo.
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Circulo representa um círculo.
type Circulo struct {
	Raio float64
}
