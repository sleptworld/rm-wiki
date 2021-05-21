package Model

import (
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/Middleware"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
)

func LoginCheck(l *Login, db *gorm.DB) (Middleware.CustomClaims,bool) {
	r := Middleware.CustomClaims{}

	if l.Email == "" || l.Pwd == ""{
		return r,false
	}
	user, _, err := DB.FindUser(db, "Email = ?", l.Email, 1)

	if err != nil {
		return r,false
	} else {
		d := (user[0]["pwd"]).([]byte)
		err := tools.PwdConfirm(l.Pwd, d, Config.AesKey)
		if err != nil {
			return r,false
		} else {
			r.ID = (user[0]["id"]).(uint)
			r.Email = (user[0]["email"]).(string)

			return r,true
		}

	}
}

func RegCheck(r *Reg, db *gorm.DB) (bool, error) {
	r_u := DB.User{
		Name:        r.Name,
		Email:       r.Email,
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
		return false, err
	}
	return true, nil
}

func EntryCheck(e *NewEntry) error {
	rE := DB.Entry{
		Title: e.Title,
		UserID: e.Author,
		Content: e.Content,
		Tags: DB.Tags2Entry(e.Tags),
		Info: e.Info,
	}

	if r,err := DB.CatCheck(e.Cat);err == nil{
		rE.CatID = r.ID
		_, err := DB.CreateEntry(DB.Db,&rE)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}

}
