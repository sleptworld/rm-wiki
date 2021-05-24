package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"github.com/sleptworld/test/tools"
	"net/http"
	"strconv"
)

func GETUser(c *gin.Context) {
	if Authentication(c,3){
		var res []DB.User
		l, p, err := common(c,DB.Db,"name",nil,&res)
		var reset []Model.AllUser
		Model.User2AllUser(&res,&reset)
		if err != nil {
			code,msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":msg,
				"Reason":code,
			},0,nil))
			return
		}
		c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,p,l,reset))
		return
	} else {
		c.JSON(http.StatusUnauthorized,Model.Api(http.StatusUnauthorized,Config.ApiVersion, map[string]string{
			"Message":Config.MsgUnauthorized,
			"Reason":Config.ErrUnauthorized,
		},0,nil))
		return
	}
}

func POSTUser(c *gin.Context){
	var regmodel Model.Reg
	if err := c.BindJSON(&regmodel);err == nil{
		var res Model.AllUser
		if l,err := Model.RegCheck(&regmodel,&res);err == nil{
			c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,int(l),res))
			return
		} else {

			code,msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":msg,
				"Reason":code,
			},0,nil))

		}
	}
}

func GETUserByID(c *gin.Context) {
	id := c.Param("id")
	var userId uint
	if uid, err := strconv.Atoi(id);err == nil{
		userId = uint(uid)
		IDorGroupCheck(c,3,userId, func(c *gin.Context) {
			var res DB.User
			h := DB.FindUser(DB.Db.Preload("Entries").Preload("EditEntries").Preload("Drafts"),
				"id = ?",id,1,&res)

			var result Model.SingleUser
			Model.User2User(&res,&result)
			c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,int(h.RowsAffected),result))

			c.Abort()
			return
		})
	} else {
		c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message":Config.MsgValueFormatForID,
			"Reason":Config.ErrValueFormat,
		},0,nil))

		return
	}

}

func PATCHUserById(c *gin.Context){
	id := c.Param("id")
	var uid uint
	if tools.StringTypeCheck(id,"Num"){
		u,_ := strconv.Atoi(id)
		uid = uint(u)
		IDorGroupCheck(c,3,uid, func(c *gin.Context) {
			var reset Model.UserUpdate
			c.BindJSON(&reset)
			r := Model.UserUpdate2User(&reset)
			user, err := DB.UpdateUser(DB.Db, "id = ?", id, 1, &r)
			if err != nil {
				c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
					"Message":Config.MsgBodyFormat,
					"Reason":Config.ErrBodyFormat,
				},0,nil))

				c.Abort()
				return
			}
			c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,int(user),reset))
		})
	}
}

func DELETEUserById(c *gin.Context){

}