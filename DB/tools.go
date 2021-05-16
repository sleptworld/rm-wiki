package DB

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/sleptworld/test/tools"
	"math/rand"
)

func isContain(items []string , item string) bool{
	for _,eachitem := range items{
		if eachitem == item{
			return true
		}
	}
	return false
}
func mapToStruct(m map[string]interface{},s interface{}) (error){
	if err := mapstructure.Decode(m,s);err != nil{
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func UserPretreatment(u *User,p string){
	rand.Read(aesKey)
	pwd,err := tools.PwdEncrypt(p,aesKey)
	if err == nil{
		u.Pwd = pwd
	}
}