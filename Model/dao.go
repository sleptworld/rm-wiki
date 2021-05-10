package Model

import (
	"errors"
	"gorm.io/gorm"
)

func search (db *gorm.DB,query string,value string,number int) ([]map[string]interface{},*gorm.DB,error){
	r := db.Where(query,value)
	if r.Error == nil{
		result := map[string]interface{}{}
		var results []map[string]interface{}
		switch number {
		case 0:
			r.Take(&result)
		case -1:
			r.Last(&result)
		case 1:
			r.First(&result)
		default:
			r.Find(&results)
			return results,r,nil
		}

		last := []map[string]interface{}{result}
		return last,r,nil
	} else {
		return nil,nil,r.Error
	}
}

func RegisterUser(db *gorm.DB,user *User) (int64,error) {
	result := db.Create(user)
	if result.Error != nil {
		return result.RowsAffected,nil
	}else {
		return 0,result.Error
	}
}

func UpdateUser(db *gorm.DB,query string,value string,number int,user *User) (int64,error) {

	_,h,err := FindUser(db,query,value,number)
	if err == nil{
		res := h.Updates(user)
		if res.Error != nil{
			return 0,res.Error
		}
		return res.RowsAffected,nil
	} else {
		return 0,nil
	}
}

func DeleteUser(db *gorm.DB, condition string, value string, number int) (int64,error) {
	ans,r,err := FindUser(db,condition,value,number)
	if ans == nil {
		return 0, nil
	}
	if err == nil{
		result := r.Delete(&User{})
		if result.Error != nil{
			return 0,result.Error
		}else {
			return result.RowsAffected,nil
		}
	} else {
		return 0,err
	}
}

func FindUser(db *gorm.DB, query string, value string, number int) ([]map[string]interface{},*gorm.DB,error){
	m := db.Model(&User{})
	rs,r,err := search(m,query,value,number)
	return rs,r,err
}

func UserForeigned(db *gorm.DB,condition string) ([]map[string]interface{},error){
	choice := []string{"Entries","EditEntries"}
	var result []map[string]interface{}
	if IsContain(choice,condition){
		r := db.Association(condition).Find(&result)
		return result,r
	} else {
		return nil,errors.New("Invalid condition.")
	}
}

func CreateEntry(db *gorm.DB, entry *Entry){
	db.Where(Entry{Title: entry.Title}).FirstOrCreate(entry)
}

func FindEntry(db *gorm.DB, condition string,value string,number int)  ([]map[string]interface{},*gorm.DB,error){
	rs,h,err := search(db.Model(&Entry{}),condition,value,number)
	return rs,h,err
}

func DeleteEntry(db *gorm.DB,condition string,value string,number int)  (int64,error){

	_,h,err := FindEntry(db,condition,value,number)
	if err == nil{
		result := h.Delete(&Entry{})
		return result.RowsAffected,nil
	} else {
		return 0,err
	}
}