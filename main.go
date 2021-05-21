package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Router"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

		DB.Db.Debug().AutoMigrate(&DB.UserGroup{},&DB.Cat{}, &DB.User{},&DB.Entry{}, &DB.History{}, &DB.Tag{})

		span := opentracing.StartSpan("gormTracint")
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(), span)

		DB.Db = DB.Db.WithContext(ctx).Debug()

		r := gin.New()

		v1 := r.Group("/v1")

		Router.EntryRouter(v1)
		Router.UserRouter(v1)
		Router.TokenRouter(v1)

		r.Run()
	}
	return
}
