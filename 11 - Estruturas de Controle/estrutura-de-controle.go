package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := -11

	if numero > 15 {
		fmt.Println("Maior do que 15")
	} else {
		fmt.Println("Menor do que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior do que zero")
	} else if numero < -10 {
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
