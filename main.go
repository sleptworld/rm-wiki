package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sleptworld/test/Model"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable")
	{
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&Model.History{}, &Model.Entry{}, &Model.User{})

			//Model.RegisterUser(db, &Model.User{
			//	Name:        "t",
			//	Email:       "t",
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


		defer func (){
			recover()
		}()
		Model.FindUser(db,"id = ?","2",1)
		db.Close()
	}