package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

// Register all routers
func Register(templates *template.Template) {

	router := mux.NewRouter()
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	router.HandleFunc("/home", hc.get)

	uic := new(uiController)
	uic.template = templates.Lookup("ui.html")
	router.HandleFunc("/ui", uic.get)

	tbc := new(tableController)
	tbc.template = templates.Lookup("table.html")
	router.HandleFunc("/table", tbc.get)

	dc := new(detailController)
	dc.template = templates.Lookup("detail.html")
	router.HandleFunc("/detail/{id:[0-9]+}", dc.get).Methods("GET")

	frc := new(formsController)
	frc.template = templates.Lookup("forms.html")
	router.HandleFunc("/forms", frc.get)

	lgc := new(loginController)
	lgc.template = templates.Lookup("login.html")
	router.HandleFunc("/login", lgc.get)

	blc := new(blankController)
	blc.template = templates.Lookup("blank.html")
	router.HandleFunc("/blank", blc.get)

	http.Handle("/", router)

	http.HandleFunc("/img/", serveResources)
	http.HandleFunc("/css/", serveResources)
	http.HandleFunc("/js/", serveResources)
	http.HandleFunc("/imp/", serveResources)
}

func serveResources(w http.ResponseWriter, r *http.Request) {
	path := "public" + r.URL.Path
	var contentType string
	if strings.HasSuffix(path, "css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)

	} else {
		w.WriteHeader(404)
	}
}
