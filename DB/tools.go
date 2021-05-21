package DB

import (
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/tools"
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

func CatCheck(c string) (Cat,error) {
	r,res := CreateCat(Db,c)

	return r,res.Error
}