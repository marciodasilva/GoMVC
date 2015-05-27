package viewmodels

import (

)

type UI struct {
	Title string
	Active string
}

func GetUI() UI{
	result := UI{
		Title: "UI Title",
		Active: "ui",
	}
	return result
}

