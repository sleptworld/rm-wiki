package Router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sleptworld/test/Controller/v1"
)

func UserRouter(group *gin.RouterGroup){
	userSource := "/user"
	group.POST(userSource,v1.RegUserHandler)
	group.GET(userSource)
	group.PATCH(userSource)
	group.DELETE(userSource)
}

func TokenRouter(group *gin.RouterGroup){
	tokenSource := "/token"
	group.POST(tokenSource)
	group.GET(tokenSource)
}
