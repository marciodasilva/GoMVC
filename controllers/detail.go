package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"wag/viewmodels"

	"github.com/gorilla/mux"
)

type detailController struct {
	template *template.Template
}

func (dc *detailController) get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idRaw := params["id"]
	id, err := strconv.Atoi(idRaw)
	log.Println("Id = ", id)
	if err == nil {
		vm := viewmodels.GetDetail(id)
		w.Header().Set("Content-Type", "text/html")
		log.Println("vm = ", vm)
		log.Println("template = ", dc.template)
		dc.template.Execute(w, vm)
		//this.template.ExecuteTemplate(w, "detail.html", vm)
	} else {
		http.Redirect(w, r, "/", 404)
	}

}
