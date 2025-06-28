package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ //Primeiro String é a Formato da Chave segundo é o valor da Chave
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Elisa",
			"sobrenome": "Neagu",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Unisa",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

}
