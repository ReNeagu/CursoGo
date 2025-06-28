package main

import "fmt"

func diaDaSemana(numero int) string { //Modo 1
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido!"
	}
}

func diaDaSemana2(numero int) string { //Modo 2, pior modo na minha opnião
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		//fallthrough Seleciona a proxima opção após o resultado
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
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Incorreto!"
	}

	return diaDaSemana
}

func diaDaSemana3(numero int) string { //Modo 3, o que eu mais gostei
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Incorreto!"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(8)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(5)
	fmt.Println(dia3)

}
