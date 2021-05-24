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
	e := Model.NewEntry{}
	err := c.BindJSON(&e)

	if e.Title == ""{
		c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message":Config.MsgBodyValueMissing,
			"Reason":Config.ErrBodyValueMissing,
		},0,nil))

		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message":Config.ErrBodyFormat,
			"Reason":Config.MsgBodyFormat,
		}, 0, nil))
		return
	}

	originResult := DB.Entry{}
	if err := Model.EntryCheck(&e, &originResult); err == nil {
		result := Model.SingleEntry{}
		Model.Entry2Entry(&originResult, &result)
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
	if !tools.StringTypeCheck(id, "Num") {
		c.JSON(http.StatusBadRequest, Model.Api(
			http.StatusBadRequest, Config.ApiVersion,
			map[string]string{
				"Message": Config.MsgValueFormatForID,
				"Reason":  Config.ErrValueFormat,
			}, 0, nil))
		return
	}
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
}

func DELETEEntryByID(c *gin.Context) {
	id := c.Param("ID")
	_, err := DB.DeleteEntry(DB.Db, "id = ?", id, 1)
	if err != nil {
		code,msg := DB.CheckErrors(err)
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message":msg,
			"Reason":code,
		},0,nil))
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Have deleted.",
		})
	}
}

func PUTEntryByID(c *gin.Context) {

}

func PATCHEntryByID(c *gin.Context) {

}
