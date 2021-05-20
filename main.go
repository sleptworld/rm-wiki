package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	v1 "github.com/sleptworld/test/Controller/v1"
	"github.com/sleptworld/test/DB"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	err error
)
func main() {
	dsn := "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable"
	DB.InitJeager()
	Database,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	{
		if err != nil {
			panic(err)
		}

		_ = Database.Use(&DB.OpentracingPlugin{})

		Database.Debug().AutoMigrate(&DB.UserGroup{},&DB.Cat{}, &DB.User{},&DB.Entry{}, &DB.History{}, &DB.Tag{}, &DB.Draft{})

		span := opentracing.StartSpan("gormTracint")
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(), span)

		session := Database.WithContext(ctx).Debug()

		r := gin.Default()

		r.POST("/reg",)

	}
	return
}
