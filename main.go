package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()

	cfg := Load_Config()

	r.HandleFunc("/HealthHandler", HealthHandler)
	r.HandleFunc("/SubmitTaskHandler", SubmitTaskHandler)
	r.HandleFunc("/EchoRequestHandler", EchoRequestHandler)
	r.HandleFunc("/SlowHandler", SlowHandler)
	r.HandleFunc("/GetTaskStatusHandler", GetTaskStatusHandler)

	r.Use(LogMiddleware, RecoverMiddleware)

	Load_DB(cfg.DBPath)

	Server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Println("[Success]server start at port", Load_Config().Port)

	if err := Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
