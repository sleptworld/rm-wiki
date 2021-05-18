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