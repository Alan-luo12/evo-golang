package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func writejson(w http.ResponseWriter, status int, code int, msg string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("json encode error", err)
		return
	}
}

func Success(w http.ResponseWriter, data any) {
	writejson(w, http.StatusOK, 0, "Success", data)
}

func Errorresponse(w http.ResponseWriter, status int, code int, msg string) {
	writejson(w, status, code, msg, nil)
}
