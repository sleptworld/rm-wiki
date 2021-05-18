package Model

import (
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"time"
)

type Login struct {
	Email string
	Pwd   string
}

type Reg struct {
	Name      string
	Email     string
	Birthday  time.Time
	Pwd       string
	Country   string
	Language  string
	Sex       int8
	Profesion string
}

func LoginCheck(l *Login, db *gorm.DB) bool {
	user, _, err := DB.FindUser(db, "Email = ?", l.Email, 1)
	if err != nil {
		return false
	} else {
		d := (user[0]["pwd"]).([]byte)
		err := tools.PwdConfirm(l.Pwd, d, Config.AesKey)
		if err != nil {
			return false
		} else {
			return true
		}

	}
}

func RegCheck(r *Reg, db *gorm.DB) bool {
	r_u := DB.User{
		Name:        r.Name,
		Email:       r.Email,
		Birthday:    r.Birthday,
		UserGroupID: 1,
		Country:     r.Country,
		Language:    r.Language,
		Entries:     nil,
		EditEntries: nil,
		Sex:         r.Sex,
		Profession:  r.Profesion,
	}
	DB.UserPretreatment(&r_u, r.Pwd)

	_, err := DB.RegisterUser(db, &r_u)
	if err != nil {
		return false
	}
	return true
}
