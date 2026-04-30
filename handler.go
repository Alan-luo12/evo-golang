package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

//返回time和status

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writejson(w, map[string]any{
		"time":   time.Now().Format(time.RFC3339),
		"status": http.StatusOK,
	})
}

// 用来接受message还有panic的结构体
type EchoRequest struct {
	Message string `json:"message"`
	Panic   bool   `json:"panic"`
}

// 返回message顺带测试panic
func EchoRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req EchoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	if req.Panic {
		panic("Manual Panic Triggered")
	}

	writejson(w, map[string]any{
		"message": req.Message,
	})
}

// 延迟函数默认100ms
func SlowHandler(w http.ResponseWriter, r *http.Request) {
	msstr := r.URL.Query().Get("ms")

	ms, err := strconv.Atoi(msstr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	time.Sleep(time.Duration(ms) * time.Millisecond)
	writejson(w, map[string]any{
		"delay_time": ms,
		"status":     "Ok",
	})

}
