package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	variavel2++
	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REF DE MEMORIA
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //DESFERENCIAÇÃO

	

}
