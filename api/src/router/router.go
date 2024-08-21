package router

import "github.com/gorilla/mux"

// Gerar vai retonar um routre com rotas devidamente configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}