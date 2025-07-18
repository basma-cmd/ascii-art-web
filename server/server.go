package server

import (
	"log"
	"net/http"

	"asciiartweb/config"
	"asciiartweb/routes"
)

// func to start our server and can handle the config
func StartServer() {
	mux := http.NewServeMux()
	routes.RoutesHandle(mux)
	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}
	log.Println("Server is running at http://localhost" + config.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
