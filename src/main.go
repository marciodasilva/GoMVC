package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"wag/controllers"
)

func main() {
	templates := populateTemplates()
	controllers.Register(templates)

	log.Println("Listening on port :8000")
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)

	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			log.Println("File Name = ", pathInfo.Name())

			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}
	result.ParseFiles(*templatePaths...)

	return result
}
