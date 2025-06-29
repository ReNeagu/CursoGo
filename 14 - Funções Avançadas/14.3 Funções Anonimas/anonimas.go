package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido --> %s", texto) //Com Sprintf da para concatenar informações
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
