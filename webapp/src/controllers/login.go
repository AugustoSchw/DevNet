package controllers

import "net/http"

// CarregarTelaDeLogin carrega a página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
