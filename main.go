package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Router"
)

func main() {

	_, err := DB.GetDbInstance().InitDBPool()

	if err == nil{
		err := DB.InitDB()
		if err != nil {
			return
		}
		r := gin.New()
		v1 := r.Group("v1")

		Router.EntryRouter(v1)
		Router.UserRouter(v1)
		Router.TokenRouter(v1)

		r.Run()
	}

	return
}
