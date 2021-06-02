package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"net/http"
)

type e struct {
	Status int
}

func GETEntry(c *gin.Context) {

	var res []Model.AllEntry
	l, p, err := common(c, DB.Db.Model(&DB.Entry{}), "Title", []string{}, &res)

	if err != nil {
		code, mes := DB.CheckErrors(err)
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message": mes,
			"Reason":  code,
		}, 0, nil))
		return
	} else {
		for step, r := range res {
			var tags []Model.Tag
			m := DB.Entry{Model: gorm.Model{ID: r.ID}}
			DB.Db.Model(&m).Association("Tags").Find(&tags)
			res[step].Tags = tags
		}
		c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, p, l, res))
	}
}

func POSTEntry(c *gin.Context) {

	e := Model.NewEntry{}
	err := c.BindJSON(&e)

	if err != nil {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message": Config.ErrBodyFormat,
			"Reason":  Config.MsgBodyFormat,
		}, 0, nil))
		return
	} else {
		var ores DB.Entry
		if err := Model.EntryCheck(&e,&ores);err == nil{
			show := Model.Entry2Entry(&ores)
			c.JSON(http.StatusOK, Model.Api(http.StatusOK,Config.ApiVersion,nil,1,show))
		} else {
			code, msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
				"Message": msg,
				"Reason":  code,
			}, 0, nil))
		}
	}
}

func GETEntryByID(c *gin.Context) {

	id := c.Param("id")

	CheckParamType(c, id, "Num", func(c *gin.Context) {
		var res DB.Entry
		r := DB.Db.Preload("Tags").Preload("History").Model(&DB.Entry{}).Where("id = ?",id).First(&res)
		if r.Error == nil {
			result := Model.Entry2Entry(&res)
			c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, nil, int(r.RowsAffected), result))
		} else {
			errCode, errMessage := DB.CheckErrors(r.Error)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
				"Message": errMessage,
				"Reason":  errCode,
			}, 0, nil))
			return
		}
	})
}

func DELETEEntryByID(c *gin.Context) {
	id := c.Param("id")
	CheckParamType(c, id, "Num", func(c *gin.Context) {

		entryID,_ := tools.String2uint(id)
		l,err := DB.Init(&DB.Entry{Model : gorm.Model{ID: entryID}}).Delete("",nil)
		if err == nil{
			c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, nil, int(l),map[string]string{
				"msg":"deleted",
			}))
			return
		} else {
			code, msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
				"Message": msg,
				"Reason":  code,
			}, 0, nil))
			return
		}
	})
}

func PATCHEntryByID(c *gin.Context) {

	id := c.Param("id")
	up := Model.UpdateEntry{}
	err := c.BindJSON(&up)
	if err != nil {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message": Config.MsgBodyFormat,
			"Reason":  Config.ErrBodyFormat,
		}, 0, nil))
		return
	}
	CheckParamType(c,id,"Num", func(c *gin.Context) {

		var ores DB.Entry
		update, err := DB.Init(&up).Update(&ores)
		if err != nil {
			code,msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusInternalServerError, Config.ApiVersion, map[string]string{
				"Message": code,
				"Reason":  msg,
			}, 0, nil))
			return
		} else {
			res := Model.Entry2Entry(&ores)
			c.JSON(http.StatusOK, Model.Api(http.StatusOK,Config.ApiVersion,nil,int(update),res))
			return
		}
	})
}
