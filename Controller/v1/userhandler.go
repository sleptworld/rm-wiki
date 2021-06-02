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

func GETUser(c *gin.Context) {
	var result []DB.User
	l,p,err := common(c,DB.Db.Model(&DB.User{}),"id",nil,&result)
	if err != nil{
		code,msg := DB.CheckErrors(err)
		c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message":msg,
			"Reason":code,
		},0,nil))
		return
	}

	res := Model.Users2Users(&result)
	c.JSON(http.StatusOK, Model.Api(http.StatusOK,Config.ApiVersion,p,int(l),res))
	return
}

func POSTUser(c *gin.Context){
	var regModel Model.Reg
	if err := c.BindJSON(&regModel);err == nil{
		var res Model.AllUser
		if l,err := Model.RegCheck(&regModel,&res);err == nil{
			c.JSON(http.StatusOK, Model.Api(http.StatusOK,Config.ApiVersion,nil,int(l),res))
			return
		} else {

			code,msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":msg,
				"Reason":code,
			},0,nil))
		}
	}
}

func GETUserByID(c *gin.Context) {
	id := c.Param("id")
	CheckParamType(c,id,"Num", func(c *gin.Context) {
		var u DB.User
		if res :=DB.Db.Preload("Entries").Preload("EditEntries").Preload("Drafts").Where("id = ?",id).First(&u);
		res.Error != nil{
			code ,msg := DB.CheckErrors(res.Error)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":msg,
				"Reason":code,
			},0,nil))
			return
		} else {
			resShow := Model.User2User(&u)
			c.JSON(http.StatusOK, Model.Api(http.StatusOK,Config.ApiVersion,nil,1,resShow))
			return
		}
	})
}

func PATCHUserById(c *gin.Context){
	id := c.Param("id")
	CheckParamType(c,id,"Num", func(c *gin.Context) {
		var updateData Model.UserUpdate
		if err := c.BindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
				"Message": "",
				"Reason":  "",
			}, 0, nil))
			return
		}
		userModel := Model.UserModel{}
		userModel.InitModel(&updateData)
		var ores DB.User
		if l, err := userModel.Update(&ores); err != nil {
			code, msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest, Config.ApiVersion, map[string]string{
				"Message": msg,
				"Reason":  code,
			}, 0, nil))
			return
		} else {
			res := Model.User2User(&ores)
			c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, nil, int(l), res))
		}
	})
}

func DELETEUserById(c *gin.Context){
	id := c.Param("id")
	CheckParamType(c,id,"Num", func(c *gin.Context) {
		uid ,_ := tools.String2uint(id)
		m := DB.Init(&DB.User{Model:gorm.Model{ID: uid}})
		if l,err := m.Delete("",nil);err != nil{
			code,msg := DB.CheckErrors(err)
			c.JSON(http.StatusBadRequest, Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":msg,
				"Reason":code,
			},0,nil))
				return
			} else {
				c.JSON(http.StatusOK, Model.Api(http.StatusOK,Config.ApiVersion,nil,int(l),map[string]string{
					"msg":"ok",}))
			}
})
}