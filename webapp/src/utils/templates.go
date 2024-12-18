package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates carrega os templates HTML na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}


// ExecutarTemplate renderiza um template html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
