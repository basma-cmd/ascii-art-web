package server

import (
	"log"
	"net/http"

	"artweb/routes"
)

func Server() {
	http.HandleFunc("/", routes.ArtWebHandler)
	http.HandleFunc("/artstyle", routes.ArtStyleHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
