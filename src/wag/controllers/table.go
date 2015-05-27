package controllers

import (
	"net/http"
	"text/template"
	"wag/viewmodels"
)

type tableController struct {
	template *template.Template
}

func (tc *tableController) get(w http.ResponseWriter, r *http.Request) {

	vm := viewmodels.GetTables()
	w.Header().Add("Content Type", "text/html")
	tc.template.Execute(w, vm)
}
