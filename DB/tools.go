package DB

import (
	"github.com/pkg/errors"
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

func CatCheck(c string,res interface{}) (int64,error) {

	num, err := Init(&Cat{Path: c}).Create(res)

	return num,err
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