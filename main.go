package main

import (
	"log"
	"os"

	"artweb/server"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("too many arguments!")
	}
	server.Server()
}
