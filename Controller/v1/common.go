package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Middleware"
	"github.com/sleptworld/test/Model"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type au func(c *gin.Context)

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

func Authentication(c *gin.Context,level int8) bool{

	cl,exist := c.Get("claims")

	if exist{

		customClaims := cl.(*Middleware.CustomClaims)

		id := customClaims.ID

		userId := strconv.Itoa(int(id))
		email := customClaims.Email

		type g struct {
			UserGroupID uint
		}

		type ll struct {
			Level int8
		}

		var u g
		var l ll

		DB.FindUser(DB.Db,"id = ? AND email = "+ "'"+email+"'",userId,1, &u)
		DB.Db.Model(&DB.UserGroup{}).Where("id = ?",u.UserGroupID).First(&l)

		if l.Level >= level{
			return true
		}
	}

	return 0 > level
}

func CheckByID(c *gin.Context,id uint) bool  {
	var claims *Middleware.CustomClaims
	if cl,exist := c.Get("claims");exist{
		claims = cl.(*Middleware.CustomClaims)

		if claims.ID == id{
			return true
		}
	}

	return false

}

func IDorGroupCheck(c *gin.Context,level int8,id uint,handler au){
	if Authentication(c,level) || CheckByID(c,id) {
		handler(c)
	} else {
		c.JSON(http.StatusUnauthorized,Model.Api(http.StatusUnauthorized,Config.ApiVersion, map[string]string{
			"Message":Config.ErrUnauthorized,
			"Reason":Config.ErrUnauthorized,
		},0,nil))
		return
	}
}

func CheckParamType(c *gin.Context,p string,t string,h au)  {
	if tools.StringTypeCheck(p,t){
		h(c)
	} else {
		c.JSON(http.StatusBadRequest,Model.Api(http.StatusBadRequest,Config.ApiVersion, map[string]string{
			"Message":Config.MsgValueFormatForID,
			"Reason":Config.ErrValueFormat,
		},0,nil))
	}
}

func Claims2Level(c *gin.Context) (int8,uint) {

	cl,exist := c.Get("claims")
	if !exist{
		return 0,1
	}

	claims := cl.(*Middleware.CustomClaims)
	var user DB.User
	var group DB.UserGroup

	r := DB.FindUser(DB.Db,"id = ?",(*claims).ID,1,&user)
	if r.Error != nil{
		return  0,1
	}
	DB.Db.Model(&DB.UserGroup{}).Select("Level").Where("id = ?",user.UserGroupID).First(&group)
	return group.Level,(*claims).ID
}