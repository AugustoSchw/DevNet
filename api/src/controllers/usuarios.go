package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CriarUsuario abre a conexão com o banco de dados para criar um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro) // Todos os tratamentos de erro em "usuarios.go" são temporarios, e serão tratados de forma adequeda no futuro
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil { // Transforma a request em uma variavel do tipo Usuario
		log.Fatal(erro)
	}

	db, erro := banco.Conectar() // Abre a conexão com o banco de dados
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close() //	No fim da função, fecha a conexão com o banco de dados

	repositorio := repositorios.NovoRepositorioDeUsuarios(db) // Adiciona o banco no repositorio
	usuarioID, erro := repositorio.Criar(usuario)             // Lida com a lógica de adicionar um usuário no banco
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuários..."))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário..."))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário..."))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário..."))
}
