package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 检查健康，返回时间
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	Success(w, map[string]any{
		"time": time.Now().Format(time.RFC3339),
	})
}

// 响应panic以及将接受到的数据反序列化之后把messgae字段发送回去
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

// 测试延迟接口，模拟慢响应真实业务
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

// 三态流转，提交任务
type SubmitRequest struct {
	Name       string
	Delay_time int
}

func SubmitTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Errorresponse(w, 405, 405, "Method Not Allowed")
		return
	}

	var req SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Errorresponse(w, 400, 400, "Bad Request"+err.Error())
		return
	}
	delay_time := req.Delay_time

	res, err := DB.Exec("INSERT INTO tasks (name,status,delay_time) VALUES(?,'pending',?)", req.Name, req.Delay_time)
	if err != nil {
		Errorresponse(w, 500, 500, "Internal Serever Error"+err.Error())
		log.Println("Create task Error", err.Error())
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		Errorresponse(w, 500, 500, "Internal Server Error"+err.Error())
		log.Println("failed to get id", err)
		return
	}

	go Processtask(id, delay_time)

	Success(w, map[string]any{
		"task_id": id,
		"status":  "submitted",
	})
}

// 查询当前的任务状态
func GetTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Errorresponse(w, 400, 400, "Method Not Allowed")
		return
	}
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		Errorresponse(w, 405, 405, "BadRequest"+err.Error())
		return
	}

	var status, result sql.NullString
	err = DB.QueryRow("SELECT status,result FROM tasks WHERE id = ?", id).Scan(&status, &result)
	if err != nil {
		if err == sql.ErrNoRows {
			Errorresponse(w, 404, 404, "task not founded"+err.Error())
			return
		}

		Errorresponse(w, 500, 500, "Internal Server Error"+err.Error())
		log.Println("failed to Queryrow ", err)
		return
	}

	Success(w, map[string]any{
		"status": status,
		"result": result,
		"id":     id,
	})
}
