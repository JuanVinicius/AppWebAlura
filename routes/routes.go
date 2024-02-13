package routes

import (
	"net/http"

	"github.com/juanvinicius/AppWebAlura/controllers"
)

func TraceRoutes() {
	http.HandleFunc("/", controllers.Index)
}
