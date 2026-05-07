package response

import (
	"app/pkg/errors"
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, code int, msg string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[Error]failed to encode error :%v", err)
		return
	}
}

func Success(w http.ResponseWriter, data any) {
	WriteJSON(w, http.StatusOK, 0, "Success", data)
}

func Error(w http.ResponseWriter, err error) {
	if apperr, ok := err.(*errors.AppError); ok {
		log.Printf("[Error] statsu:%v  code:%v  msg:%v  err:%v", apperr.HTTPStatus(), apperr.Code, apperr.Msg, apperr.Err)
		WriteJSON(w, apperr.HTTPStatus(), apperr.Code, apperr.Msg, nil)
		return
	}

	log.Printf("[Error] statsu:%v  code:%v  msg:%v  err:%v", http.StatusInternalServerError, 5000, "InrternalServerError", err)
	WriteJSON(w, http.StatusInternalServerError, 5000, "InrternalServerError", nil)
}
