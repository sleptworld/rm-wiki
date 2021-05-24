package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func common(c *gin.Context,tx *gorm.DB,order string,preloads []string,res interface{}) (int,map[string]string,error) {

	l := c.DefaultQuery("limit","10")
	of := c.DefaultQuery("offset","-1")
	or := c.DefaultQuery("order",order)
	la := c.DefaultQuery("lang","zh")

	limit,err1 := strconv.Atoi(l)
	offset,err2 := strconv.Atoi(of)

	if err1 == nil && err2 == nil {
		if limit < 0 || offset < -1 {
			return 0,nil,errors.New("wrong")
		}

		temp := tx.Limit(limit).Offset(offset).Order(or)

		for _,preload := range preloads{
			temp = temp.Preload(preload)
		}

		rsss := temp.Find(res)

		if rsss.Error != nil{
			return 0,nil,rsss.Error
		}

		p := map[string]string{
			"limit" : l,
			"offset" : of,
			"order" : or,
			"lang" : la,
		}

		return int(rsss.RowsAffected),p,nil

	} else {
		return 0,nil,errors.New("wrong format")
	}
}
