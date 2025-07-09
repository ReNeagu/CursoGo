package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //Driver Do PostgreSQL
)

func main() {
	host := "localhost"
	porta := 5432
	usuario := "renato"
	senha := "renato1234"
	banco := "testdb"

	urlConexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, porta, usuario, senha, banco)

	db, err := sql.Open("postgres", urlConexao)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao testar conexão:", err)
	}

	fmt.Println("✅ Conectado com sucesso ao PostgreSQL!")

}
