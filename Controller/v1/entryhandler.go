package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"net/http"
	"strconv"
)

type e struct {
	Status int
}

func GETEntry(c *gin.Context) {
	limit := c.DefaultQuery("limit","10")
	offset := c.Query("offset")
	order := c.DefaultQuery("order","Title")

	var res []DB.Entry
	bad := gin.H{
		"400":gin.H{
			"error_code": "20001",
			"error": "bad request",
		},
	}

	l,err1 := strconv.Atoi(limit)
	o,err2 := strconv.Atoi(offset)

	if err1 == nil && err2 == nil {
		if l <= 0 || l > 20 || o < -1 {
			c.JSON(http.StatusOK, bad)
			return
		}

		DB.Db.Limit(l).Offset(o).Order(order).Find(&res)
		c.JSON(http.StatusOK,gin.H{
			"200":gin.H{
				"data" : res,
			},
		})

		return

	} else {
		c.JSON(http.StatusOK,gin.H{
			"400":gin.H{
				"error_code":"20002",
				"error":"limit and offset must be interger.",
			},
		})

		return
	}
}