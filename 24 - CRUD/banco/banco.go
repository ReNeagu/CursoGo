package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver de Conexão com MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	usuario := "neagu"
	senha := "ne@gu18"
	host := "localhost"
	porta := "3306"
	banco := "testdb"

	stringConexao := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		usuario, senha, host, porta, banco)

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
