package main

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"net/http"
)

type TemplateData struct {
	URL       string
	Flash     string
	Error     string
	CSRFToken string
}

// defaultData used for some defalt data
func (a *application) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.URL = a.server.url

	return td
}

func (a *application) render(w http.ResponseWriter, r *http.Request, view string, vars jet.VarMap) error {
	td := &TemplateData{}

	td = a.defaultData(td, r)

	tp, err := a.view.GetTemplate(fmt.Sprintf("%s.html", view))
	if err != nil {
		return err
	}

	if err = tp.Execute(w, vars, td); err != nil {
		return err
	}

	return nil
}
