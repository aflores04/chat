package utils

import (
	renderPkg "github.com/unrolled/render"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HttpResponse(w http.ResponseWriter, code int, body interface{}) {
	rend := renderPkg.New()

	_ = rend.JSON(w, code, body)
}

func HttpErrorResponse(w http.ResponseWriter, code int, err error) {
	rend := renderPkg.New()
	resp := ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}

	_ = rend.JSON(w, code, resp)
}
