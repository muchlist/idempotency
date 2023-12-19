package main

import (
	"idempotency/module/book"
	"idempotency/pkg/mlogger"
	"net/http"

	"idempotency/pkg/lrucache"
	"idempotency/pkg/middleware"

	"github.com/go-chi/chi/v5"
)

func mapRoutes(log mlogger.Logger) http.Handler {
	r := chi.NewRouter()

	// dependency
	cache := lrucache.NewLRUCache()

	// middleware
	idempo := middleware.NewIdempotencyMiddleware(cache)

	bookRepo := book.NewRepo()
	bookUsecase := book.NewUsecase(bookRepo)
	bookHandler := book.NewHandler(bookUsecase, log)

	// Endpoint with auth

	r.Route("/books", func(r chi.Router) {
		// not using custom idempotency
		r.Get("/{id}", bookHandler.Get)

		// using middleware idempotency
		i := r.With(idempo.IdempotentCheck)
		i.Post("/", bookHandler.Insert)
		i.Put("/{id}", bookHandler.Update)
	})

	return r
}
