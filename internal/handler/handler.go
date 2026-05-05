package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"app/internal/model"
	"app/internal/service"
	"app/pkg/errors"
	"app/pkg/response"
)

//三个简单接口函数

type TaskHandler struct {
	svc *service.TaskService
}

func (h *TaskHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	response.Success(w, map[string]string{
		"time": time.Now().Format(time.RFC3339),
	})
}

type EchoRequest struct {
	Message string `json:"message"`
	Panic   bool   `json:"panic"`
}

func (h *TaskHandler) EchoRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, errors.NewUserError(40051, "Method Not Allowed", nil))
		return
	}
	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, errors.NewUserError(4007, "bad request", err))
		return
	}
	if req.Panic {
		panic("manual Panic triggered")
	}
	response.Success(w, req.Message)
}

func (h *TaskHandler) SlowHandler(w http.ResponseWriter, r *http.Request) {
	msstr := r.URL.Query().Get("ms")
	ms, err := strconv.Atoi(msstr)
	if err != nil {
		ms = 100
	}
	time.Sleep(time.Duration(ms) * time.Millisecond)
	response.Success(w, map[string]int{
		"delay_time": ms,
	})
}

//用于DI注入，分层解耦思想

func NewTaskHandler(svc *service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

//带数据库的业务接口
//对于service层返回出来的错误直接用response.Error()即可，但是对于本层自己出现的错误需要再()里面自己再写进去errorrs.NewXxxError

func (h *TaskHandler) Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, errors.NewUserError(4003, "invalid json", nil))
		return
	}

	var t model.SubmitTaskReq
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		response.Error(w, errors.NewUserError(4004, "Bad Request", err))
		return
	}

	resp, err := h.svc.SubmitTask(t)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, resp)
}

func (h *TaskHandler) Getstatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, errors.NewUserError(4005, "invalid method", nil))
		return
	}

	idstr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		response.Error(w, errors.NewUserError(4006, "Bad Request", err))
		return
	}

	resp, err := h.svc.GetTaskStatus(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, resp)
}
