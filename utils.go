package main

import (
	"encoding/json"
	"net/http"
)

func writejson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
