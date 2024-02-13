package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
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

	db := connDB()
	defer db.Close()

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom", 59, 1},
		{"PN", "show", 99, 11},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}

func connDB() *sql.DB {
	conn := "user=postgres dbname=postgres password=Abc@123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
