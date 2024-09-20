package main

import (
	"log"
	"net/http"
)

func (a *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	err := a.render(w, r, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}
