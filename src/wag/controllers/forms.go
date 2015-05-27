package controllers

import (
	"net/http"
	"text/template"
	"wag/viewmodels"
)

type formsController struct {
	template *template.Template
}

func (this *formsController) get(w http.ResponseWriter, r *http.Request) {
	vmf := viewmodels.GetForms()
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vmf)
}
