package main

import (
	"net/http"
)

func main() {
	// Create mux that routes URLs to the correct handler.
	mux := http.NewServeMux()

	// Create a web server that runs on port 8080.
	newServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	// Serve files from the current folder for every URL.
	// mux.Handle("/app/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	})

	// Start the server and wait for requests.
	newServer.ListenAndServe()

}
