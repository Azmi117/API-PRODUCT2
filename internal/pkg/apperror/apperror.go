package apperror

import "net/http"

// 1. Buat struct Apperror

type Apperror struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// 2. Buat method error agar termasuk kedalam error interface

func (e *Apperror) Error() string {
	return e.Message
}

// 3. Buat function-function custom error (NotFound, BadRequest, dll)
func NotFound(msg string) error {
	return &Apperror{Code: http.StatusNotFound, Message: msg}
}

func BadRequest(msg string) error {
	return &Apperror{Code: http.StatusBadRequest, Message: msg}
}

func Forbidden(msg string) error {
	return &Apperror{Code: http.StatusForbidden, Message: msg}
}

func UnAuthorized(msg string) error {
	return &Apperror{Code: http.StatusUnauthorized, Message: msg}
}

func Internal(msg string) error {
	return &Apperror{Code: http.StatusInternalServerError, Message: msg}
}
