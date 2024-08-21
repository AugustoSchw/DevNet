package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa a estrutura geral de todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar retorna o router já com todas as rotas configuradas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas { // Configura cada rota
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
