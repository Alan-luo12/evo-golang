package response

import (
	"Lin/pkg/errors"
	"encoding/json"
	"log"
	"net/http"
)

// Response结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 写入JSON响应
func WriteJSON(w http.ResponseWriter, status int, code int, msg string, data any) {
	//设置响应头
	w.Header().Set("Content-Type", "application/json")
	//写入响应头
	w.WriteHeader(status)

	//创建响应体
	resp := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	//编码响应体
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[Error]failed to encode error :%v", err)
		return
	}
}

// 成功响应
func Success(w http.ResponseWriter, data any) {
	WriteJSON(w, http.StatusOK, 0, "Success", data)
}

// 错误响应
func Error(w http.ResponseWriter, err error) {
	//如果err是AppError类型，则返回对应的HTTP状态码和错误信息
	if apperr, ok := err.(*errors.AppError); ok {
		log.Printf("[Error] statsu:%v  code:%v  msg:%v  err:%v", apperr.HTTPStatus(), apperr.Code, apperr.Msg, apperr.Err)
		WriteJSON(w, apperr.HTTPStatus(), apperr.Code, apperr.Msg, nil)
		return
	}
	//如果err不是AppError类型，则返回500错误
	log.Printf("[Error] statsu:%v  code:%v  msg:%v  err:%v", http.StatusInternalServerError, 5000, "InrternalServerError", err)
	WriteJSON(w, http.StatusInternalServerError, 5000, "InrternalServerError", nil)
}
