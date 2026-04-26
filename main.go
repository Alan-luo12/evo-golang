package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/status", Status)
	log.Println("The Programe will be running on the port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Sorry error in starting port", err)
		return
	}
}
