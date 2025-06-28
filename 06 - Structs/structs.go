package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

func main() {
	fmt.Println("Teste")

	var u usuario
	u.nome = "Renato"
	u.idade = 31
	fmt.Println(u)

	enderecoExemplo := endereco{"José Luis Monteiro", 825}

	usuario2 := usuario{"Karen", 35, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Lili"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 1}
	fmt.Println(usuario4)

}
