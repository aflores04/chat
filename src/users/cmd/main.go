package main

import (
	"github.com/aflores04/chat/src/mongodb"
	"github.com/aflores04/chat/src/users/db"
	"github.com/aflores04/chat/src/users/handler"
	"github.com/aflores04/chat/src/users/service"
	"github.com/alecthomas/inject"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Start(handler handler.UserHandler) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})

	http.ListenAndServe(":3001", r)
}

func main() {
	injector := inject.New()
	injector.Install(
		&mongodb.MongoModule{},
		&db.UserRepositoryModule{},
		&service.UserServiceModule{},
		&handler.UserHandlerModule{},
	)
	injector.Call(Start)
}
