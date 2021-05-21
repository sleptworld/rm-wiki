package Router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sleptworld/test/Controller/v1"
)

func UserRouter(group *gin.RouterGroup){
	userSource := "/user"

	group.GET(userSource,v1.GETUser)
	group.POST(userSource,v1.POSTUser)

	uGroup := group.Group(userSource)
	{
		uGroup.GET("/:id",v1.GETUserByID)
		uGroup.PATCH("/:id",v1.PATCHUserById)
		uGroup.DELETE("/:id",v1.DELETEUserById)
	}
}