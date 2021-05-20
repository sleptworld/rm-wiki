package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"net/http"
)

func UserHandler(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func RegUserHandler(c *gin.Context){
	var regmodel Model.Reg
	if c.BindJSON(&regmodel) == nil{
		if r,err := Model.RegCheck(&regmodel,DB.Db);r{
			c.JSON(http.StatusOK,gin.H{
				"status" : 1,
				"msg" : "ok",
				"data" : "ok",
			})
		} else {
			c.JSON(http.StatusOK,gin.H{
				"status" : -1,
				"msg": err.Error(),
			})
		}
	}
}