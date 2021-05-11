package main

import (
	"fmt"
	"github.com/sleptworld/test/Model"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
)

func t (a interface{}) (interface{}){
	typeOfA := reflect.TypeOf(a)
	alns := reflect.New(typeOfA)
	fmt.Println(alns.Type())
	return alns
}

func main() {
	dsn := "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	{
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&Model.UserGroup{}, &Model.User{}, &Model.Entry{}, &Model.History{})

		//Model.CreateGroup(db,[]Model.UserGroup{
		//	{
		//		GroupName: "Admin",
		//		Users: nil,
		//		Level: 3,
		//	},
		//	{
		//		GroupName: "Walker",
		//		Users: nil,
		//		Level: 0,
		//	},
		//})

		//Model.RegisterUser(db,&Model.User{
		//	Name:        "ruomu",
		//	Email:       "test@test.com",
		//	Pwd:         "t",
		//	UserGroupID: 1,
		//	Avatar:      "",
		//	Description: "",
		//	Site:        "",
		//	Country:     "",
		//	Language:    "",
		//	Entries:     []Model.Entry{
		//		{
		//			Title: "no",
		//			History: nil,
		//			Content: "hello",
		//		},
		//	},
		//	EditEntries: nil,
		//	Mechanism:   "",
		//	Sex:         0,
		//	Profession:  "",
		//})

	}

	//Model.CreateEntry(db,&Model.Entry{
	//	Title:   "test4",
	//	UserID:  4,
	//	Tags:    "",
	//	Cat:     "",
	//	History: nil,
	//	Content: "test4",
	//})

	Model.UpdateUser(db,"id = ?","3",1,&Model.User{
		Email:       "ruomu@gmail.com",
	})
	return
}
