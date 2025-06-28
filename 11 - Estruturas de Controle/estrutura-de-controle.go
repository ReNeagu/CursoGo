package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 20

	if numero > 15 {
		fmt.Println("Maior do que 15")
	} else {
		fmt.Println("Menor do que 15")
	}

	if outroNumero := numero; outroNumero > 30 {
		fmt.Println("Número é maior do que 0")
	} else {
		fmt.Println("Numero é menor que 30")
	}

}
