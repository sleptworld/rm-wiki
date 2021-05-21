package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"net/http"
	"strconv"
)

type e struct {
	Status int
}

func GETEntry(c *gin.Context) {
	limit := c.DefaultQuery("limit","10")
	offset := c.DefaultQuery("offset","-1")
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
			c.JSON(http.StatusBadRequest, bad)
			return
		}

		DB.Db.Limit(l).Offset(o).Order(order).Preload("History").Preload("Tags").Find(&res)
		c.JSON(http.StatusOK,gin.H{
			"200":gin.H{
				"data" : res,
			},
		})

		return

	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"400":gin.H{
				"error_code":"20002",
				"error":"limit and offset must be interger.",
			},
		})

		return
	}
}

func POSTEntry(c *gin.Context){
	e := Model.NewEntry{}
	c.BindJSON(&e)

	if err := Model.EntryCheck(&e);err == nil{
		c.JSON(http.StatusOK,gin.H{
			"200":gin.H{
				"msg":"OK",
			},
		})
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"400":gin.H{
				"msg": err.Error(),
			},
		})
	}
}

func GETEntryByID(c *gin.Context){
	id := c.Param("ID")

	e, _, err := DB.FindEntry(DB.Db, "id = ?", id, 1)

	if err == nil{
		c.JSON(http.StatusOK,gin.H{
			"200":gin.H{
				"data": e[0],
			},
		})
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"400":err.Error(),
		})
	}
}

func DELETEEntryByID(c *gin.Context)  {
	id := c.Param("ID")
	_, err := DB.DeleteEntry(DB.Db,"id = ?",id,1)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"400":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"200": "delete",
		})
	}
}

func PUTEntryByID(c *gin.Context)  {

}

func PATCHEntryByID(c *gin.Context){

}
