package controllers

import (
	"net/http"
	"text/template"
	"wag/viewmodels"
)

type uiController struct {
	template *template.Template
}

func (this *uiController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetUI()
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}
