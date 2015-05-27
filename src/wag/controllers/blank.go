package controllers

import (
	"net/http"
	"text/template"
	"wag/viewmodels"
)

type blankController struct {
	template *template.Template
}

func (this *blankController) get(w http.ResponseWriter, r *http.Request) {
	vmb := viewmodels.GetBlank()
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vmb)
}
