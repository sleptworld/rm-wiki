package main

import (
	"github.com/sleptworld/test/Model"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	{
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&Model.User{}, &Model.Entry{}, &Model.History{})

		//Model.RegisterUser(db, &Model.User{
		//	Name:        "ruomu",
		//	ID:          "test@test.com",
		//	Pwd:         "t",
		//	Avatar:      "t",
		//	Description: "t",
		//	Site:        "t",
		//	Country:     "t",
		//	Language:    "t",
		//	Mechanism:   " ",
		//	Sex:         1,
		//	Profession:  "t",
		//})
	}

	n,_,_ := Model.FindUser(db,"name = ?","ruomu",1)
	if n ==nil{}
}
