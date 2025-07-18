package main

import (
	"asciiartweb/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("No Arguments needed (:")
		return
	}
	// start server 
	server.StartServer()
}