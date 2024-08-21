package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retonar um routre com rotas devidamente configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()	// Cria router
	return rotas.Configurar(r)	// Devolve o router, já configurado
}