package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sleptworld/test/Model"
)

func main(){
	db,err := gorm.Open("postgres","host=localhost user=postgres dbname=postgres password=1234567 sslmode=disable")
	{
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&Model.Entry{},&Model.User{},&Model.Comments{})

		Model.CreateNewUser("A"," "," "," "," "," "," ",true,db)
	}
	defer db.Close()
}
