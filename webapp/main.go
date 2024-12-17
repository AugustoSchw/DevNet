package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	// Start the web server
	fmt.Println("webapp iniciado")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
