package Router

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup){
	uGroup := group.Group("/user")
	{
		lGroup := uGroup.Group("/login")
		{
			lGroup.GET("/:name")
			lGroup.PATCH("/:name")
		}
		uGroup.POST("/register")
	}
}
