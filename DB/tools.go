package DB

import (
	"errors"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
)



func UserPretreatment(u *User,p string){
	pwd,err := tools.PwdEncrypt(p,Config.AesKey)
	if err == nil{
		u.Pwd = pwd
	}
}

func Tags2Entry(t []string) []Tag{
	var res []Tag
	for _,name := range t{
		res = append(res, Tag{
			Name:    name,
		})
	}

	return res
}

func CatCheck(c string,res interface{}) *gorm.DB {
	r := CreateCat(Db,c,res)

	return r
}

func CheckErrors(err error) (code string,msg string){
	if errors.Is(err,gorm.ErrRecordNotFound){
		return Config.ErrNoValue,Config.MsgNoValueForID
	}

	if errors.Is(err,gorm.ErrModelValueRequired) {
		return Config.ErrBodyValueMissing,Config.MsgBodyValueMissing
	}

	if errors.Is(err,gorm.ErrRegistered){
		return Config.ErrBodyRegistered,Config.MsgBodyValueRegistered
	}

	if errors.Is(err,gorm.ErrInvalidData) {
		return Config.ErrWrongData,Config.MsgWrongData
	}

	return "",""
}