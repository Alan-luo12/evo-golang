package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if !AllowOnlyGet(w, r) {
		return
	}

	name := r.URL.Query().Get("name")
	if name != "laura" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "Welcome %s", name)
	log.Printf(" %s %s ?name=%s\n", r.Method, r.URL.Path, name)
}

func Status(w http.ResponseWriter, r *http.Request) {
	AllowOnlyGet(w, r)
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "now [%s] the Server is running", now)
	log.Printf("%s %s", r.Method, r.URL.Path)
}
