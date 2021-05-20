package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Model"
	"gorm.io/gorm"
	"net/http"
)

func UserHandler(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func RegUserHandler(db *gorm.DB,c *gin.Context){
	var regmodel Model.Reg
	if c.BindJSON(&regmodel) == nil{
		if Model.RegCheck(&regmodel,db){
			c.JSON(http.StatusOK,gin.H{
				"status" : 1,
				"msg" : "ok",
				"data" : "ok",
			})
		} else {
			c.JSON(http.StatusOK,gin.H{
				"status" : -1,
				"msg": "wrong",
			})
		}
	}
}