package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()

	return mux
}
