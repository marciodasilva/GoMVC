package controllers

import (
	"net/http"
	"text/template"
	"wag/viewmodels"
)

type loginController struct {
	template *template.Template
}

func (this *loginController) get(w http.ResponseWriter, r *http.Request) {
	vml := viewmodels.GetLogin()
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vml)
}
