package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Middleware"
	"gorm.io/driver/postgres"
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
	DB.InitJeager()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	{
		if err != nil {
			panic(err)
		}

		_ = db.Use(&DB.OpentracingPlugin{})

		db.AutoMigrate(&DB.UserGroup{}, &DB.User{}, &DB.Entry{}, &DB.History{})

		span := opentracing.StartSpan("gormTracint")
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(),span)

		session := db.WithContext(ctx)

		//DB.CreateGroup(session,[]DB.UserGroup{
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
		//
		//DB.RegisterUser(session,&DB.User{
		//	Name:        "ruomu",
		//	Email:       "test@test.com",
		//	Pwd:         "t",
		//	UserGroupID: 1,
		//	Avatar:      "",
		//	Description: "",
		//	Site:        "",
		//	Country:     "",
		//	Language:    "",
		//	Entries:     []DB.Entry{
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

		//DB.RegisterUser(session,&DB.User{
		//	Name:        "ruomu",
		//	Email:       "ruomu@test.com",
		//	Pwd:         "t",
		//	UserGroupID: 1,
		//	Avatar:      "t",
		//	Description: "t",
		//	Site:        "t",
		//	Country:     "t",
		//	Language:    "t",
		//	Entries:     nil,
		//	EditEntries: nil,
		//	Mechanism:   "",
		//	Sex:         0,
		//	Profession:  "t",
		//})
		err := Middleware.InitSearcher(session)
		if err != nil {
			return
		} else {
			res := Middleware.Search("hello",&Middleware.SearchOpt)
			fmt.Println(res)
		}

	}

	return
}
