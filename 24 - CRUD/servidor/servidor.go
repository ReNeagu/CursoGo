package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!!"))
		return
	}
	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
	}
	defer db.Close()

	//PREPARE STATEMENT
	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	//Teste com outro formato (201)
	idInserido, err := insercao.LastInsertId()
	if err != nil {
		http.Error(w, "Erro ao obter o ID inserido: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	// Define o status 201 Created antes de escrever no body
	w.WriteHeader(http.StatusCreated)

	// Agora escreva a mensagem usando o ID
	w.Write([]byte(
		fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido),
	))

	// idInserido, erro := insercao.LastInsertId()
	// if erro != nil {
	// 	w.Write([]byte("Erro ao obter o ID Inserido!"))
	// 	return
	// }
	// //STATUS CODES

	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}
