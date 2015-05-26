package viewmodels

type Details struct {
	Title  string
	Active string
}

type Detail struct {
	Title  string
	Active string
	Detail []Users
}

type Users struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
}

func GetDetail(Id int) Detail {
	result := Detail{
		Title:  "Title [Detail]",
		Active: "detail",
	}
	row1 := Users{
		Id:        1,
		FirstName: "Marcio",
		LastName:  "DaSilva",
		UserName:  "@mdas",
	}
	row2 := Users{
		Id:        2,
		FirstName: "Jacob",
		LastName:  "Thornton",
		UserName:  "@fat",
	}
	row3 := Users{
		Id:        3,
		FirstName: "Larry",
		LastName:  "the Bird",
		UserName:  "@twitter",
	}
	result.Detail = []Users{row1, row2, row3}

	for _, v := range result.Detail {
		if v.Id == Id {
			row := Users{v.Id, v.FirstName, v.LastName, v.UserName}
			result.Detail = []Users{row}
		}
	}
	return result
}
