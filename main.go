package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sleptworld/test/DB"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type test interface {
	say()
}

type t struct {
	r int
}


func (test *t) say(){
	test.r = 10
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
		ttt := DB.User{
			Name:        "ruomu",
			Email:       "test@test.com",
			UserGroupID: 1,
			Entries:     []DB.Entry{
				{
					Title: "no",
					History: nil,
					Content: "hello",
				},
			},
			EditEntries: nil,
			Sex:         0,
		}

		DB.UserPretreatment(&ttt,"helloworld@123")
		if _,err := DB.RegisterUser(session,&ttt);err != nil{
			fmt.Println(err)
		}


		//DB.RegisterUser(session,&DB.User{
		//	Name:        "zwp",
		//	Email:       "zwp@test.com",
		//	Pwd:         "zhouweiping123",
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
		//err := Middleware.InitSearcher(session)
		//if err != nil {
		//	return
		//} else {
		//	res := Middleware.Search("hello",&Middleware.SearchOpt)
		//	fmt.Println(res)
		//}

	}

	//b := t{r: 6}
	//b.say()
	//fmt.Println(b)



	return
}
