package main

import "fmt"

func main() {
	resultado := diaDaSemana(1)

	fmt.Println(resultado)

	fmt.Println("------------------")

	resultado2 := diaDaSemana2(1)

	fmt.Println(resultado2)
}

// Switch
// Não precisa de break, por padrão já conta com o break implícito
// Não precisa de parênteses
// Basicamente um if else simplificado
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}

}

// Outra maneira de fazer
func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		// clausula fallthrough
		// faz com que o switch continue executando os cases abaixo
		fallthrough // Não é muito utilizado
		// Basicamente passa para o próximo case direto para o próximo case
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero invalido"
	}

	return diaDaSemana
}

// Jeito mais óbvio de fazer (simplificado)
/*
func diaDaSemana(numero int) string {
	if numero == 1 {
		return "Domingo"
	}else if numero = 2 {
		return "Segunda"
	}else if numero = 3 {
		return "Terça"
	}else if numero = 4 {
		return "Quarta"
	}else if numero = 5 {
		return "Quinta"
	}else if numero = 6 {
		return "Sexta"
	}else if numero = 7 {
		return "Sabado"
	}else {
		return "Numero invalido"
	}
}
*/
