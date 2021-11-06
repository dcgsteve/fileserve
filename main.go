package main

import (
	"net/http"
)

func main() {
	// Simple static webserver:
	http.Handle("/", http.FileServer(http.Dir("./")))
}
