package Router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sleptworld/test/Controller/v1"
)

func TokenRouter(group *gin.RouterGroup) {
	tokenSource := "/Token"
	group.POST(tokenSource,v1.POSTToken)
	group.GET(tokenSource,v1.GETToken)
	group.DELETE(tokenSource,v1.DELETEToken)
}
