package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			context.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "no token",
				"data":   nil,
			})
			context.Abort()
			return
		}
	}
}
