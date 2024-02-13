package main

import (
	"net/http"

	"github.com/juanvinicius/AppWebAlura/routes"
)

func main() {
	routes.TraceRoutes()
	http.ListenAndServe(":8000", nil)
}
