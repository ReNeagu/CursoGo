package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a função init") //Ela sempre roda antes de qq função, independente de onde esteja
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada!")
	fmt.Println(n)
}
