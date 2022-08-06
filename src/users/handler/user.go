package handler

import (
	"encoding/json"
	"github.com/aflores04/chat/src/users/errors"
	"github.com/aflores04/chat/src/users/request"
	"github.com/aflores04/chat/src/utils"
	"net/http"
)

func (h userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterUserRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	resp, err := h.service.RegisterUser(r.Context(), &req)
	if err != nil {
		utils.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.HttpResponse(w, http.StatusOK, resp)
}

func (h userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	logged := h.service.LoginAttempt(r.Context(), &req)
	if !logged {
		utils.HttpErrorResponse(w, http.StatusBadRequest, &errors.InvalidUsernameOrPasswordError{})
		return
	}

	utils.HttpResponse(w, http.StatusOK, "logged in!")
}
