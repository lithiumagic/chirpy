package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	newServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	newServer.ListenAndServe()
}
