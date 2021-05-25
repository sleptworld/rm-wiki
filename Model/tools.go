package Model

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"net/http"
)

func DbError(c *gin.Context,err error){
	code,msg := DB.CheckErrors(err)
	c.JSON(http.StatusBadRequest,Api(http.StatusBadRequest,Config.ApiVersion,map[string]string{
		"Message":msg,
		"Reason":code,
	},0,nil))

	return

}
