package Web

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Router"
)

func MainRouter(e *gin.Engine)  {
	v1Group := e.Group("v1")
	Router.UserRouter(v1Group)
	Router.EntryRouter(v1Group)
}
