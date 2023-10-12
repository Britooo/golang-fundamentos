package main

import (
	"fmt"
	"math"
)

type formaGeomtrica interface {
	area() float64
}

// Exemplo de "polimorfismo"
// Podemos passar qualquer struct que implemente a interface formaGeomtrica
func area(forma formaGeomtrica) float64 {
	return forma.area()
}

// Para implementar uma interface, basta implementar todos os métodos
// Usando um receiver do tipo da struct
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2) //(c.raio * c.raio)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {

	fmt.Println("Interfaces")

	// Interfaces são tipos que definem um comportamento
	// Uma interface é um conjunto de métodos
	// Nesse caso se quisermos calcular de uma forma geomtrica
	// Não faria sentido ter um método calcularArea para cada struct
	// Por isso criamos uma interface com esse método

	retangulo1 := retangulo{10, 15}

	circulo1 := circulo{10}

	fmt.Printf("A área do retângulo é %.2f\n", area(retangulo1))

	fmt.Printf("A área do círculo é %.2f\n", area(circulo1))
}
