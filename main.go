package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/HealthHandler", HealthHandler)
	mux.HandleFunc("/EchoRequestHandler", EchoRequestHandler)
	mux.HandleFunc("/SlowHandler", SlowHandler)

	handler := LogMiddleware(RecoverMiddleware(mux))
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
