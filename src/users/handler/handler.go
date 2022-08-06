package handler

import (
	"github.com/aflores04/chat/src/users/service"
	"net/http"
)

type UserHandlerModule struct{}

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service service.UserService
}

func (*UserHandlerModule) ProvideUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}
