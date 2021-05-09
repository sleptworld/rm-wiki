package Model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var availableQueryUser = []string{"id","name","sex","country","language","email","profession"}

func search (db *gorm.DB,query string,value string,number int) ([]map[string]interface{},*gorm.DB,error){
	r := db.Where(query,value)
	if r.Error == nil{
		var result map[string]interface{}
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
		return nil,nil,errors.New("wrong")
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

func UpdateUser(db *gorm.DB,user *User) (error) {
	result := db.Model(&User{}).Updates(user)
	if result.Error != nil{
		return nil
	}
	return nil
}

//func DeleteUser(db *gorm.DB, condition string, value string, number int) (int64,error) {
//	ans,r,err := FindUser(db,condition,value,number)
//	if ans == nil {
//		return 0, nil
//	}
//	if err == nil{
//		result := r.Delete(&User{})
//		if result.Error != nil{
//			return 0,result.Error
//		}else {
//			return result.RowsAffected,nil
//		}
//	} else {
//		return 0,errors.New("Wrong")
//	}
//}

func FindUser(db *gorm.DB, query string, value string, number int) {
	//r := db.Where(query,value)

	var test map[string]interface{}
	db.Model(&User{}).First(&test,"id = ?","1")
	if test == nil{}
	//rs,r,err := search(m,query,value,number)

	//if rs == nil && r == nil && err == nil{}
	//if r.Error == nil{
	//	var result User
	//	var results []User
	//	switch number {
	//	case 0:
	//		r.Take(&result)
	//	case -1:
	//		r.Last(&result)
	//	case 1:
	//		r.First(&result)
	//	default:
	//		r.Find(&results)
	//		return results,r,nil
	//	}
	//	last := []User{result}
	//	return last,r,nil
	//}else {
	//	return nil,nil,r.Error
	//}
}

func CreateEntry(db *gorm.DB, entry *Entry){
	db.FirstOrCreate(entry,entry)
}

func FindEntry(db *gorm.DB, condition string,value string)  {
}

func DeleteEntry(db *gorm.DB,)  {
}