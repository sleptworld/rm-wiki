package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Middleware"
	"github.com/sleptworld/test/Router"
)

func main() {

	_,err := DB.GetDbInstance().InitDBPool()
	if err == nil{
		err := DB.InitDB()
		if err != nil {
			return
		}

		e := gin.New()

		e.Use(gin.Logger(),gin.Recovery(),Middleware.Jwt(),Middleware.Permission())

		DB.AddRole([][]string{
			{"Anonymous","/v1/Entry/*","GET"},
			{"SuperAdmin","/v1/*","*"},
		})

		DB.AddRoleForUser(0,"Anonymous")
		v1Group := e.Group("v1")
		{
			Router.EntryRouter(v1Group)
		}

		e.Run()

	}
	return
}
