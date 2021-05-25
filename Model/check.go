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
	var user DB.User
	h := DB.FindUser(db, "Email = ?", l.Email, 1,&user)

	if h.Error != nil {
		return r,false
	} else {
		d := user.Pwd
		err := tools.PwdConfirm(l.Pwd, d, Config.AesKey)
		if err != nil {
			return r,false
		} else {
			r.ID = user.ID
			r.Email = user.Email

			return r,true
		}

	}
}

func RegCheck(r *Reg,res interface{})  (int64,error) {
	r_u := DB.User{
		Name:        r.Name,
		Email:       r.Email,
		UserGroupID: 2,
		Country:     r.Country,
		Language:    r.Language,
		Sex:         r.Sex,
		Profession:  r.Profession,
	}
	DB.UserPretreatment(&r_u, r.Pwd)

	l, err := DB.RegisterUser(DB.Db, &r_u,res)
	return l,err
}

func EntryCheck(e *NewEntry,id uint,rev bool,res interface{}) error {
	rE := DB.Entry{
		Title: e.Title,
		UserID: id,
		Content: e.Content,
		Tags: DB.Tags2Entry(e.Tags),
		Review: rev,
		Info: e.Info,
	}


	result := Cat{}
	if r := DB.CatCheck(e.Cat,&result);r.Error == nil{
		rE.CatID = result.ID
		_, err := DB.CreateEntry(DB.Db,&rE,res)
		if err != nil {
			return err
		}
		return nil
	} else {
		return r.Error
	}

}

func UpdateEntryCheck(e *UpdateEntry,id uint,res interface{}) error{
	rE := DB.Entry{
		Model:gorm.Model{ID:id},
		Content: e.Content,
		Tags: DB.Tags2Entry(e.Tags),
	}
	err := DB.UpdateEntry(DB.Db,&rE,res)
	return err
}
