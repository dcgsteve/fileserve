package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// Get user specified info
	envFileDir := os.Getenv("FS_DIR")
	if envFileDir == "" {
		envFileDir = "./"
	}

	envPort := os.Getenv("FS_PORT")
	if envPort == "" {
		envPort = "8080"
	}

	// Serv files
	http.Handle("/", http.FileServer(http.Dir(envFileDir)))
	log.Printf("Listening on port %s\n", envPort)
	err := http.ListenAndServe(":"+envPort, nil)
	if err != nil {
		log.Fatal(err)
	}

}
