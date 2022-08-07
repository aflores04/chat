package main

import (
	"github.com/aflores04/chat/src/jwt"
	"github.com/aflores04/chat/src/mongodb"
	"github.com/aflores04/chat/src/users/db"
	"github.com/aflores04/chat/src/users/handler"
	"github.com/aflores04/chat/src/users/service"
	"github.com/alecthomas/inject"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func Start(
	handler handler.UserHandler,
) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", "http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Origin",
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"Referer",
			"Referrer-Policy"},
	}))

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})

	http.ListenAndServe(":3001", r)
}

func main() {
	injector := inject.New()
	injector.Install(
		&mongodb.MongoModule{},
		&jwt.JwtModule{},
		&db.UserRepositoryModule{},
		&service.UserServiceModule{},
		&handler.UserHandlerModule{},
	)
	injector.Call(Start)
}
