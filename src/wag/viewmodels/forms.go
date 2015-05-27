package viewmodels

import (

)

type Forms struct {
	Title string
	Active string
}

func GetForms() Forms{
	result := Forms{
		Title: " Zontal ",
		Active: "forms",
	}
	return result
}

