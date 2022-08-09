package main

import (
	"github.com/aflores04/chat/backend/src/chat_messages/db"
	"github.com/aflores04/chat/backend/src/chat_messages/handler"
	"github.com/aflores04/chat/backend/src/chat_messages/service"
	"github.com/aflores04/chat/backend/src/mongodb"
	"github.com/alecthomas/inject"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func Start(
	handler handler.ChatMessagesHandler,
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

	r.Route("/messages", func(r chi.Router) {
		r.Get("/list", handler.GetMessages)
	})

	log.Println("Starting messages service ... ")
	http.ListenAndServe(":3002", r)
}

func main() {
	injector := inject.New()
	injector.Install(
		&mongodb.MongoModule{},
		&db.ChatMessagesRepositoryModule{},
		&service.ChatMessagesServiceModule{},
		&handler.ChatMessagesHandlerModule{},
	)
	injector.Call(Start)
}
