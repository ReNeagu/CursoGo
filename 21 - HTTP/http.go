package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar Página de Usuários!"))
}

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação web

	//CLIENTE --> SERVIDOR

	//ROTAS
	// URI - Identificador do Recurso
	// MÉTODO - GET (buscar), POST (cadastrar), PUT (atualizar), DELETE

	//...............................................................
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Olá Mundo!"))

	// })

	// http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Carregar Página de Usuários!"))

	// })
	//.............................................................

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil)) //aqui criamos o servidor local
}
