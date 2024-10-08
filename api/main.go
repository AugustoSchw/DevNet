package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	r := router.Gerar()                                                 // Gera router
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)) // Setta o nosso "site" com localhost:5000/
}
