package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"gorm.io/gorm"
	"net/http"
)

func readEntries(c *gin.Context){

	var res []Model.AllEntry
	tmpDB := DB.Db.Model(&DB.Entry{})
	l, p, err := common(c, tmpDB, "Title", []string{}, &res)
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
		a := Model.Api(http.StatusOK, Config.ApiVersion, p, l, res)
		c.JSON(http.StatusOK, a)
	}
}

func readEntry(c *gin.Context,e *DB.Entry){
	var Res Model.SingleEntry
	Model.Entry2Entry(e,&Res)
	c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,1,Res))
}

func writeEntry(c *gin.Context,id uint) {
	var e Model.NewEntry
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message": Config.ErrBodyFormat,
			"Reason":  Config.MsgBodyFormat,
		}, 0, nil))
		return
	}

	var review bool
	if id == Config.AnonymousID{
		review = true
	} else {
		review = false
	}

	var originRes DB.Entry
	if err := Model.EntryCheck(&e,id,review,&originRes);err != nil{
		code,msg := DB.CheckErrors(err)

		c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message" : msg,
			"Reason":code,
		},0,nil))
		return
	} else {
		var res Model.SingleEntry
		Model.Entry2Entry(&originRes,&res)
		c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,1,res))
		return
	}
}

func deleteEntry(c *gin.Context,e *DB.Entry){

	l,err := DB.DeleteEntry(DB.Db,"id = ?",e.ID)
	if err != nil{
		Model.DbError(c,err)
		return
	}

	c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,int(l),map[string]string{
		"msg":"OK",
	}))

	return
}

func updateEntry(c *gin.Context,e *DB.Entry){

	var uData Model.UpdateEntry
	err := c.BindJSON(&uData)
	if err != nil {
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
			"Message": Config.ErrBodyFormat,
			"Reason":  Config.MsgBodyFormat,
		}, 0, nil))
		return
	}

	var oRes DB.Entry
	if err := Model.UpdateEntryCheck(&uData,e.ID,&oRes);err != nil{
		Model.DbError(c,err)
		return
	}

	var Res Model.SingleEntry
	Model.Entry2Entry(&oRes,&Res)
	c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,1,Res))
	return
}

// User

func readUsers(c *gin.Context){

	var originUsers DB.User
	l,p,err := common(c,DB.Db,"id",[]string{"Drafts","EditEntries","Entries"},&originUsers)

	if err != nil{
		c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion,p,l,nil))
		return
	}


}