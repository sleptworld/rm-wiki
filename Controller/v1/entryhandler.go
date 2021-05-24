package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type e struct {
	Status int
}

func GETEntry(c *gin.Context) {

	var res []Model.AllEntry

	tmpDB := DB.Db.Model(&DB.Entry{})

	l, p, err := common(c, tmpDB, "Title", []string{}, &res)

	if err != nil {

		code,mes := DB.CheckErrors(err)
		c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message":mes,
			"Reason":code,
		},0,nil))

		return

	} else {
		for step, r := range res {
			var tags []Model.Tag
			m := DB.Entry{Model: gorm.Model{ID: r.ID}}
			DB.Db.Model(&m).Association("Tags").Find(&tags)
			res[step].Tags = tags
		}
		a := Model.Api(http.StatusOK, Config.ApiVersion, p, l, res)
		c.JSON(http.StatusOK, a)
	}
}

func POSTEntry(c *gin.Context) {

	l,uid := Claims2Level(c)

	e := Model.NewEntry{}
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message":Config.ErrBodyFormat,
			"Reason":Config.MsgBodyFormat,
		}, 0, nil))
		return
	}

	var oresult DB.Entry

	var policy bool

	switch l {
	case 3:
		policy = false
	default:
		policy = true
	}

	if err := Model.EntryCheck(&e,uid,policy,&oresult);err == nil{
		result := Model.SingleEntry{}
		Model.Entry2Entry(&oresult, &result)
		c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, nil, 1, result))
	} else {
		code,msg := DB.CheckErrors(err)
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message":msg,
			"Reason":code,
		}, 0, nil))
	}
}

func GETEntryByID(c *gin.Context) {
	id := c.Param("id")

	CheckParamType(c,id,"Num", func(c *gin.Context) {
		var res DB.Entry
		r := DB.FindEntry(DB.Db.Preload("Tags").Preload("History"),
			"id = ?", id, 1, &res)

		if r.Error == nil {
			result := Model.SingleEntry{}
			Model.Entry2Entry(&res, &result)
			c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, nil, int(r.RowsAffected), result))
		} else {
			errCode,errMessage := DB.CheckErrors(r.Error)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":errMessage,
				"Reason":errCode,
			},0,nil))
			return
		}
	})
}

func DELETEEntryByID(c *gin.Context) {
	id := c.Param("id")

	CheckParamType(c,id,"Num", func(c *gin.Context) {

		var entry DB.Entry
		if r:= DB.FindEntry(DB.Db,"id = ?",id,1,&entry);r.Error==nil{
			IDorGroupCheck(c,3,entry.UserID, func(c *gin.Context) {
				if l,err := DB.DeleteEntry(DB.Db,"id = ?",id,1);err == nil{
					c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,int(l),entry))
					return
				} else {
					code,msg := DB.CheckErrors(err)
					c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
						"Message":msg,
						"Reason":code,
					},0,nil))
					return
				}
			})
		} else {
			c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message" : Config.MsgNoValueForID,
				"Reason" : Config.ErrValueFormat,
			},0,nil))
		}
	})
}

func PATCHEntryByID(c *gin.Context) {
	l,uid := Claims2Level(c)
	
	i := strconv.Itoa(int(uid))
	IDorGroupCheck(c,3,i, func(c *gin.Context) {
		up := Model.UpdateEntry{}
		err := c.BindJSON(&up)
		if err != nil {
			return
		}
		res := Model.EntryUpdate2Entry(&up)
		DB.UpdateEntry(DB.Db,&res,uid)
	})

}
