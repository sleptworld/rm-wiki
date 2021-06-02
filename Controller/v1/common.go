package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type au func(c *gin.Context)
type ac func(c *gin.Context, id uint, level DB.Level)

// tool functions

func common(c *gin.Context, tx *gorm.DB, order string, preloads []string, res interface{}) (int, map[string]string, error) {

	l := c.DefaultQuery("limit", "10")
	of := c.DefaultQuery("offset", "-1")
	or := c.DefaultQuery("order", order)
	la := c.DefaultQuery("lang", "zh")

	limit, err1 := strconv.Atoi(l)
	offset, err2 := strconv.Atoi(of)

	if err1 == nil && err2 == nil {
		if limit < 0 || offset < -1 {
			return 0, nil, errors.New("wrong")
		}

		temp := tx.Limit(limit).Offset(offset).Order(or)
		for _, preload := range preloads {
			temp = temp.Preload(preload)
		}
		rsss := temp.Find(res)
		if rsss.Error != nil {
			return 0, nil, rsss.Error
		}

		p := map[string]string{
			"limit":  l,
			"offset": of,
			"order":  or,
			"lang":   la,
		}
		return int(rsss.RowsAffected), p, nil
	} else {
		return 0, nil, errors.New("wrong format")
	}
}


func CheckParamType(c *gin.Context,id string,t string,handler au)  {
	if tools.StringTypeCheck(id,t){
		handler(c)
	} else {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion,map[string]string{
			"msg":"a",
			"c":"d",
		},1,nil))
	}
}