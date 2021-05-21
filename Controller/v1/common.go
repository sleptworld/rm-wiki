package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func common(c *gin.Context,key string,order string){

	l := c.DefaultQuery("limit","10")
	of := c.DefaultQuery("offset","1")
	order := c.DefaultQuery("order",order)

	limit,err1 := strconv.Atoi(l)
	offset,err2 := strconv.Atoi(of)

	if err1 == nil && err2 == nil {
		if limit < 0 || offset < -1 {
			return false
		}

	} else {
		return false
	}
}
