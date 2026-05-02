package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	r.HandleFunc("/HealthHandler", HealthHandler)
	r.HandleFunc("/EchoRequestHandler", EchoRequestHandler)
	r.HandleFunc("/SlowHandler", SlowHandler)

	r.Use(LogMiddleware, RecoverMiddleware)

	server := http.Server{
		Addr:    ":" + Load_Config().Port,
		Handler: r,
	}

	log.Println("the programe wiil be running on the port :", Load_Config().Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
