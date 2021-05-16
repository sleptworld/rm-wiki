package Model

import (
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
)

type Login struct {
	Email string
	Pwd string
}

func LoginCheck(l *Login,db *gorm.DB){
	user, _, err := DB.FindUser(db, "Email = ?", l.Email, 1)
	if err != nil {
		return
	} else {
		d := user[0]["pwd"]

		tools.PwdConfirm(l.Pwd,d,[]byte{})
	}
}