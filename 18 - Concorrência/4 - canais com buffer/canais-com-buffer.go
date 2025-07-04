package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"

	mensagem := <-canal //aqui a seta indica que o conteúdo que entrou em canal esta indo para mensagem
	fmt.Println(mensagem)
}
