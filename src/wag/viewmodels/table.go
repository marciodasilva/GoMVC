package viewmodels

//"log"

type Tables struct {
	Title  string
	Active string
	Users  []User
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
}

func GetTables() Tables {
	result := Tables{
		Title:  "Title [Table]",
		Active: "table",
	}
	row1 := User{
		Id:        1,
		FirstName: "Marcio",
		LastName:  "DaSilva",
		UserName:  "@mdas",
	}
	row2 := User{
		Id:        2,
		FirstName: "Jacob",
		LastName:  "Thornton",
		UserName:  "@fat",
	}
	row3 := User{
		Id:        3,
		FirstName: "Larry",
		LastName:  "the Bird",
		UserName:  "@twitter",
	}
	result.Users = []User{row1, row2, row3}

	return result
}

// func GetTable(Id int) Tables {
// tables := GetTables()
//	newTables := Tables{
//		Title:  "Title [Table]",
//		Active: "table",
//	}
//	for _, v := range tables.Users {
//		if v.Id == Id {
//			row := User{v.Id, v.FirstName, v.LastName, v.UserName}
//			newTables.Users = []User{row}
//		}
//	}
//	return newTables
//}
