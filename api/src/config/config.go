package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com MySQL
	StringConexaoBanco = ""
	// Porta é onde a API vai estar rodando no momento
	Porta = 0
)

// Carregar vai inicializar as variáveis de ambienet
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)		// Mata a execução, pois as informações são esseciais para o inicio do programa
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}