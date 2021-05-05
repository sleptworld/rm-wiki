package Model

import "github.com/jinzhu/gorm"

func CreateNewUser(Name string,Email string,pwd string,Site string,Country string,language string,M string,Sex bool,db *gorm.DB)  {

	simple := User{
		Name: Name,
		Email: Email,
		Pwd: pwd,
		Sex: Sex,
		Site: Site,
		Country: Country,
		Language: language,
		Mechanism: M,
	}

	db.Save(&simple)

}