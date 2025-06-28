package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 { //Modo 1
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println("Total de:", i)

	// for j := 0; j < 20; j += 2 { //Modo 2
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes { //Trocar o UNDERLINE POR UMA VARIAVEL PARA CARREGAR O INDICE
		fmt.Println(nome) //ADD A VARIAVEL DE INDICE CRIADA AQUI
	}

	//Teste para Atualização Git

}
