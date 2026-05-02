package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	Success(w, map[string]any{
		"time": time.Now().Format(time.RFC3339),
	})
}

type EchoResponse struct {
	Msg   string `json:"msg"`
	Panic bool   `json:"panic"`
}

func EchoRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Errorresponse(w, http.StatusMethodNotAllowed, 0, "Method Not Allowed")
		return
	}
	var resp EchoResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		Errorresponse(w, http.StatusInternalServerError, 0, "Invalid json")
		log.Println("json encode error", err)
		return
	}

	if resp.Panic {
		panic("manual panic triggered")
	}

	Success(w, map[string]any{
		"msg": resp.Msg,
	})
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	msstr := r.URL.Query().Get("ms")
	ms, err := strconv.Atoi(msstr)
	if err != nil {
		Errorresponse(w, http.StatusBadRequest, 400, "BadRequest")
		log.Println("Atoi error", err)
		return
	}

	time.Sleep(time.Duration(ms) * time.Millisecond)

	Success(w, map[string]any{
		"delay_time": ms,
	})
}
