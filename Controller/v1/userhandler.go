package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Model"
	"net/http"
	"strconv"
)

func GETUser(c *gin.Context) {
	p := access{
		idAccess:           false,
		MinimumPermissions: Config.AdminLevel,
		handlers : map[DB.Level]ac{
			Config.AdminLevel: func(c *gin.Context, id uint, level DB.Level) {
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
			}},
		}

		p.policy(c)

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
	CheckParamType(c,id,"Num", func(c *gin.Context) {
		var u DB.User
		var uid uint
		if res :=DB.FindUser(DB.Db.Preload("Entries").Preload("EditEntries").Preload("Drafts"),
			"id = ?",id,1,&u);res.Error != nil{
			code ,msg := DB.CheckErrors(res.Error)
			c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
				"Message":msg,
				"Reason":code,
			},0,nil))
			return
		}
		intNum,_ := strconv.Atoi(id)
		uid = uint(intNum)
		p := access{
			idAccess:           true,
			id:                 uid,
			MinimumPermissions: 3,
			handlers: map[DB.Level]ac{
				Config.AdminLevel : func(c *gin.Context, id uint, level DB.Level) {
					var result Model.SingleUser
					Model.User2User(&u,&result)
					c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,1,result))
					c.Abort()
					return
				},
			},
		}
	p.policy(c)
	})
}

func PATCHUserById(c *gin.Context){
	id := c.Param("id")
	CheckParamType(c,id,"Num", func(c *gin.Context) {
		intNum,_ := strconv.Atoi(id)
		uid := uint(intNum)
		p := access{
			idAccess:           true,
			id:                 uid,
			MinimumPermissions: Config.AdminLevel,
			handlers: map[DB.Level]ac{
				Config.AdminLevel : func(c *gin.Context, id uint, level DB.Level) {
					var updateData Model.UserUpdate
					if err := c.BindJSON(&updateData);err != nil{
						c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
							"Message":"",
							"Reason":"",
						},0,nil))
						return
					}
					reset := Model.UserUpdate2User(&updateData)

					cb := DB.User{}
					if l,err := DB.UpdateUser(DB.Db,&reset,&cb);err != nil{
						code,msg := DB.CheckErrors(err)
						c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
							"Message":msg,
							"Reason":code,
						},0,nil))
						return
					}else {
						var rb Model.SingleUser
						Model.User2User(&cb,&rb)
						c.JSON(http.StatusOK, Model.Api(http.StatusOK, Config.ApiVersion, nil, int(l),rb))
					}
				},
			},
		}
		p.policy(c)
	})
}

func DELETEUserById(c *gin.Context){
	id := c.Param("id")
	CheckParamType(c,id,"Num", func(c *gin.Context) {
		intNum,_ := strconv.Atoi(id)
		uid := uint(intNum)
		p := access{
			idAccess:           true,
			id:                 uid,
			MinimumPermissions: Config.AdminLevel,
			handlers: map[DB.Level]ac{
				Config.AdminLevel: func(c *gin.Context, i uint, level DB.Level) {
					if l,err := DB.DeleteUser(DB.Db,"id = ?",id,1);err != nil{
						code,msg := DB.CheckErrors(err)
						c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
							"Message":msg,
							"Reason":code,
						},0,nil))
						return
					} else {
						c.JSON(http.StatusOK,Model.Api(http.StatusOK,Config.ApiVersion,nil,int(l),map[string]string{
							"msg":"ok",
						}))
					}
				},
			},
		}
		p.policy(c)
	})
}