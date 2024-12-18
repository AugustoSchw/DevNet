package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	// Start the web server
	utils.CarregarTemplates()
	r := router.Gerar()
	
	
	fmt.Println("webapp iniciado")
	log.Fatal(http.ListenAndServe(":3000", r))
}
