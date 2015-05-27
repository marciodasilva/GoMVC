package viewmodels

import (

)

type Login struct{
	Title string
	Active string
}

func GetLogin() Login{
	result := Login{
		Title: "Login Title",
		Active: "login",
	}
	return result
}

