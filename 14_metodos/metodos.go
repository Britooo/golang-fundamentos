package main

import "fmt"

type usuario struct {
	primeiroNome string
	ultimoNome   string
	idade        uint8

	// Isso não é um método, é apenas um campo do tipo função, não é isso que queremos
	//salvar func()
}

// Para criar um método, basta criar uma função com um receiver
// O receiver é um parâmetro que indica em qual tipo a função será associada
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.primeiroNome)
}

// Se assemelha a um método de uma classe
func (user usuario) getFullname() string {
	return user.primeiroNome + " " + user.ultimoNome
}

func (user usuario) isMaiorDeIdade() bool {
	return user.idade >= 18
}

// Podemos ter métodos com ponteiros também
// Isso é útil quando queremos alterar o valor de um campo da struct
// Se não usarmos ponteiros, a alteração não será refletida na struct original
// Isso é semelhante a um método setter
func (user *usuario) fazerAniversario() {
	user.idade++
}

func (user usuario) teste() {
	user.primeiroNome = "Teste"
}

func main() {

	fmt.Println("Métodos")

	// Métodos são funções que agem em um tipo específico
	// e estão associadas a structs, interfaces, etc

	// Algo semelhante aos comportamentos em classes (Java, C#, etc)
	// Porém, em Go não temos classes, nem é orientado a objetos

	usuario1 := usuario{"Fulano", "de Tal", 17}

	usuario2 := usuario{"Ciclano", "primo do Beltrano", 18}

	usuario1.salvar()

	fmt.Printf("O nome completo do usuário é %s\n", usuario1.getFullname())

	fmt.Printf("É maior de idade? %t\n", usuario1.isMaiorDeIdade())

	fmt.Printf("\n")
	usuario1.fazerAniversario()
	usuario2.fazerAniversario()

	fmt.Printf("%s agora tem %d anos\n", usuario1.primeiroNome, usuario1.idade)

	fmt.Printf("%s agora tem %d anos\n", usuario2.primeiroNome, usuario2.idade)

	fmt.Printf("\n")
	fmt.Printf("%s é maior de idade? %t\n", usuario1.primeiroNome, usuario1.isMaiorDeIdade())
	fmt.Printf("%s é maior de idade? %t\n", usuario2.primeiroNome, usuario2.isMaiorDeIdade())

	// O receiver é uma cópia do valor original
	// Por isso, não é possível alterar o valor original
	usuario1.teste()

	fmt.Printf("\n")
	fmt.Println("O nome do usuário não foi alterado")
	fmt.Printf("O nome completo do usuário é %s\n", usuario1.getFullname())
}
