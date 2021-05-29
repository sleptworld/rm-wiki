package main

import (
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	err error
)

func main() {

	_, err := DB.GetDbInstance().InitDBPool()

	if err == nil{
		err := DB.InitDB()
		if err != nil {
			return
		}

		test := (&Model.UserModel{})
		test.InitModel(&Model.Reg{
			Name:       "tcc",
			Email:      "tcc@123.com",
			Pwd:        "fadaf",
			Country:    "",
			Language:   "",
			Sex:        1,
			Profession: "",
		})

		var result DB.User
		test.Create(&result)
	//	r := gin.New()
	//	r.Use(gin.Logger(),Middleware.Jwt(),gin.Recovery())
	//	v1 := r.Group("v1")
	//
	//	Router.EntryRouter(v1)
	//	Router.UserRouter(v1)
	//	Router.TokenRouter(v1)
	//
	//	r.Run()
	}

	return
}
