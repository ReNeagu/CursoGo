package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//CREATE = MÉDOTO POST
	//READ = MÉDOTO GET
	//UPDATE = MÉDOTO PUT
	//DELETE = MÉDOTO DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id})", servidor.BuscarUsuario).Methods(http.MethodGet)

	fmt.Println("Escutando na Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
