package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"net/http"
)

type AllCat struct {
	ID   int32
	Path string
}

func GETCat(c *gin.Context) {
	var res []AllCat

	_,_,err := common(c,DB.Db.Model(&DB.Cat{}),"path",[]string{},&res)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"data":res,
		})
	}

}

func POSTCat(c *gin.Context){}

func GETCatByID(c *gin.Context){

	var result []DB.Cat
	id := c.Param("id")
	pre := DB.Db.Model(&DB.Cat{}).Where("id = ?",id).Preload("Entries")
	_,_,err := common(c,pre,"path",[]string{},&result)

	if err == nil{
		c.JSON(http.StatusOK,gin.H{
			"data" : result,
		})
	} else {
		c.JSON(http.StatusBadGateway,gin.H{
			"err":err.Error(),
		})
	}
}
