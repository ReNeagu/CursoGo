package main

import "fmt"

func main() {
	//Operadores Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 12 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//Operadores de Atribuição
	var variavel1 string = "String"
	variavel := "String2"
	fmt.Println(variavel1, variavel)

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Operadores Lógicos
	fmt.Println("---------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //AND
	fmt.Println(verdadeiro || falso) //OR
	fmt.Println(!verdadeiro)         //negação, para inverter o valor

	//Operadores Unários (para incrementar o valor de uma variavel)
	numero := 10
	numero++
	numero += 15
	numero *= 2
	fmt.Println(numero)

}
