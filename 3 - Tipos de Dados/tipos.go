package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -10000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE

	var numero3 rune = 12334
	fmt.Println(numero3)

	// BYTE =  UNIT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123342312.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1231.44
	fmt.Println(numeroReal3)

	var str string = "Texto1"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'R'
	fmt.Println(char)

	var vazia int16
	fmt.Println(vazia)

	//Se não for declarado true ou false, o sistema considera False
	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
	//Retorno do erro é <nil>

	var erro2 error = errors.New("Erro Interno!")
	fmt.Println(erro2)

}
