package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{"João", "joao.silva@email.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})
	fmt.Println("Escutando na Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil)) //aqui criamos o servidor local
}
