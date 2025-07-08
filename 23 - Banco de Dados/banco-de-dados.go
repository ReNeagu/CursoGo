// package main

// func main() {
// 	//urlConexao := "root:neagu/banco?charset=utf8&parseTime=true&loc=Local"
// }

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver do MySQL
)

func main() {
	// Altere conforme os dados que você criou no Workbench
	usuario := "neagu"
	senha := "ne@gu18"
	host := "localhost"
	porta := "3306"
	banco := "testdb"

	// Monta a string de conexão no formato que o driver do MySQL espera:
	//usuario:senha@tcp(host:porta)/banco?charset=utf8&parseTime=true&loc=Local
	urlConexao := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		usuario, senha, host, porta, banco)

	// Abre a conexão com o banco
	db, err := sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}
	defer db.Close() // Garante que a conexão será fechada ao final

	// Testa se a conexão está ativa
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar conexão:", err)
	}

	fmt.Println("✅ Conectado com sucesso ao banco MySQL!")
}
