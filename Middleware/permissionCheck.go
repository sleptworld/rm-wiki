package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"net/http"
)

func Permission() gin.HandlerFunc  {
	return func(context *gin.Context) {
		oClaims , exist := context.Get("claims")
		var id uint

		if exist{
			claims := oClaims.(CustomClaims)
			id = claims.ID
		} else {
			id = 0
		}

		if DB.Check(id,context.Request.URL.Path,context.Request.Method){
			return
		} else {
			context.JSON(http.StatusUnauthorized, Model.Api(http.StatusUnauthorized, Config.ApiVersion, map[string]string{
				"Message": Config.MsgUnauthorized,
				"Reason":  Config.ErrUnauthorized,
			}, 0, nil))
			context.Abort()
			return
		}
	}
}
