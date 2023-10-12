package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Printf("Recebido -> %s", texto)
	}("Olá mundo") // Invocando a função anonima

	// Podemos retornar algum valor da função anonima
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá mundo") // Invocando a função anonima

	fmt.Println(retorno)

	/*
		No javascript, por exemplo, poderiamos fazer algo assim
		Exemplo:

		(()=>{
			console.log("Olá mundo"
		})();
	*/
}
