package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres dbname=wiki password=123456 sslmode=disable"
	DB.InitJeager()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	{
		if err != nil {
			panic(err)
		}

		_ = db.Use(&DB.OpentracingPlugin{})

		err := db.Debug().AutoMigrate(&DB.Cat{})
		if err != nil {
			return
		} else {
			db.Debug().Raw("CREATE INDEX path_gist_idx ON cats USING gist(path);" +
				"CREATE INDEX path_idx ON cats USING btree(path);")
		}

		span := opentracing.StartSpan("gormTracint")
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(),span)


		session := db.WithContext(ctx)

		Model.RegCheck(&Model.Reg{
			Name:      "zhouweiping",
			Email:     "zwp@qq.com",
			Pwd:       "heloworld@123",
			Country:   "zh",
			Language:  "zh",
			Sex:       0,
			Profesion: "t",
		},session)

	}

	return
}
