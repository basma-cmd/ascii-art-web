package routes

import (
	"net/http"

	"asciiartweb/controllers"
)

func RoutesHandle(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.HomeController)
	mux.HandleFunc("/ascii-art", controllers.AsciiArtController)
}

