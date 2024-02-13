package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom", 59, 1},
		{"PN", "show", 99, 11},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
