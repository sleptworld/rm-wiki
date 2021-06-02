package Model

//
//func LoginCheck(l *Login, db *gorm.DB) (Middleware.CustomClaims,bool) {
//
//	r := Middleware.CustomClaims{}
//
//	if l.Email == "" || l.Pwd == ""{
//		return r,false
//	}
//
//	var user DB.User
//	if res := DB.Init(&DB.User{Email: l.Email}).Query(&user,1,"",nil);res.Error != nil{
//		return r,false
//	}
//
//	d := user.Pwd
//	err := tools.PwdConfirm(l.Pwd, d, Config.AesKey)
//	if err != nil {
//			return r,false
//		} else {
//		r.ID = user.ID
//		r.Email = user.Email
//		return r,true
//	}
//}

func RegCheck(r *Reg,res interface{})  (int64,error) {

	userModel := UserModel{}
	userModel.InitModel(r)
	create, err := userModel.Create(res)

	return create,err
}

func EntryCheck(e *NewEntry,res interface{}) error {

	entryModel := EntryModel{}
	entryModel.InitModel(e)
	_, err := entryModel.Create(res)
	return err
}

func UpdateEntryCheck(e *UpdateEntry,res interface{}) error{

	entryModel := EntryModel{}
	entryModel.InitModel(e)
	_,err := entryModel.Update(res)
	return err
}
