package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	// Routes
	mux.Get("/", a.homeHandler)

	// fileServer to load assets and css
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return mux
}
