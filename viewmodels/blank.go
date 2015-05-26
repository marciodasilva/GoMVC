package viewmodels

import (

)

type Blank struct {
	Title string
	Active string
}

func GetBlank() Blank{
	result := Blank{
		Title: "Blank Title",
		Active: "blank",
	}
	return result
	
}
