package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	v1 "github.com/sleptworld/test/Controller/v1"
	"github.com/sleptworld/test/DB"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	var err error
	dsn := "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable"
	DB.InitJeager()
	DB.Db,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	{
		if err != nil {
			panic(err)
		}

		_ = DB.Db.Use(&DB.OpentracingPlugin{})

		DB.Db.Debug().AutoMigrate(&DB.UserGroup{},&DB.Cat{}, &DB.User{},&DB.Entry{}, &DB.History{}, &DB.Tag{}, &DB.Draft{})

		span := opentracing.StartSpan("gormTracint")
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(), span)

		DB.Db = DB.Db.WithContext(ctx).Debug()

		//DB.CreateEntry(DB.Db,&DB.Entry{
		//	Model:     gorm.Model{},
		//	Title:     "",
		//	UserID:    0,
		//	Lock:      false,
		//	EditingBy: 0,
		//	Tags:      nil,
		//	CatID:     0,
		//	History:   nil,
		//	Content:   "",
		//	Info:      "",
		//})

		//DB.CreateGroup(DB.Db,[]DB.UserGroup{
		//	DB.UserGroup{
		//		GroupName: "admin",
		//		Level:     3,
		//	},
		//	DB.UserGroup{
		//		GroupName: "anonymous",
		//		Users:     nil,
		//		Level:     0,
		//	},
		//})

		r := gin.Default()

		r.GET("/reg", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"hello":"world",
			})
		})
		r.POST("/reg",v1.RegUserHandler)

		r.GET("/Entry",v1.GETEntry)

		r.Run()
	}
	return
}
