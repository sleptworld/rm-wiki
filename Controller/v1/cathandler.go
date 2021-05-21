package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"net/http"
	"strconv"
)

type AllCat struct {
	ID   int32
	Path string
}

func GETCat(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "-1")
	order := c.DefaultQuery("order", "path")

	var res []AllCat

	l, err1 := strconv.Atoi(limit)
	o, err2 := strconv.Atoi(offset)

	if err1 == nil && err2 == nil {
		if l < 0 || l > 20 || o < -1 {
			return
		}
		if r := DB.Db.Model(&DB.Cat{}).Limit(l).Offset(o).Order(order).Find(&res);r.Error == nil{
			c.JSON(http.StatusOK,gin.H{
				"data" : res,
			})
		} else {
			c.JSON(http.StatusBadRequest,gin.H{
				"err":"err",
			})
		}
		} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"err",
		})
	}
}

func GETCatByID(c *gin.Context){
	id := c.Param("id")

	l := c.DefaultQuery("limit","10")
	of := c.DefaultQuery("offset","1")
	order := c.DefaultQuery("order","path")

	limit,err1 := strconv.Atoi(l)
	offset,err2 := strconv.Atoi(of)

	if err1 == nil && err2 == nil {
		if limit < 0 || offset < -1 {
			return
		}

	}
}