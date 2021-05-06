package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sleptworld/test/Model"
)

func main(){
	db,err := gorm.Open("postgres","host=localhost user=ruomu dbname=wiki password=123456 sslmode=disable")
	{
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&Model.History{},&Model.Entry{},&Model.User{})
	}
	db.Close()
}
