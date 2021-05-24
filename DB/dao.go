package DB

import (
	"errors"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"strings"
)

type Dao interface {
	Create()
	Delete()
	Update()
	Query()
}

// Tools for search

func search (db *gorm.DB,query string,value interface{},number int,res interface{}) *gorm.DB{
	var r *gorm.DB
	r = db.Where(query,value)
	if r.Error == nil{
		switch number {
		case 0:
			r.Take(res)
		case -1:
			r.Last(res)
		case 1:
			r.First(res)
		default:
			r.Find(res)
		}

		return r

	} else {
		return nil
	}
}

// User DB's Dao

func RegisterUser(db *gorm.DB,user *User,res interface{}) (int64,error) {
	// Email only
	result := db.Where(User{Email:user.Email}).FirstOrCreate(user)
	FindUser(db,"email = ?",user.Email,1,res)
	return result.RowsAffected,result.Error
}

func UpdateUser(db *gorm.DB,query string,value string,number int,user *User) (int64,error) {

	r := FindUser(db,query,value,number,&User{})
	if r.Error == nil{
		if res := r.Updates(user); res.Error != nil{
			return 0,res.Error
		} else {
			return res.RowsAffected,nil
		}
	} else {
		return 0,nil
	}
}

func DeleteUser(db *gorm.DB, condition string, value string, number int) (int64,error) {
	r := FindUser(db,condition,value,number,&User{})
	if r.Error == nil{
		result := r.Delete(&User{})
		if result.Error != nil{
			return 0,result.Error
		}else {
			return result.RowsAffected,nil
		}
	} else {
		return 0,r.Error
	}
}

func FindUser(db *gorm.DB, query string, value interface{}, number int,res interface{}) *gorm.DB{
	m := db.Model(&User{})
	result := search(m,query,value,number,res)
	return result
}

func UserForeignKey(db *gorm.DB,m map[string]interface{},condition string) ([]map[string]interface{},error){
	choice := []string{"Entries","EditEntries"}
	userModel := User{
		Model:gorm.Model{
			ID: m["id"].(uint),
		},
	}
	var result []map[string]interface{}
	if tools.IsContain(choice,condition){
		err := db.Model(&userModel).Association(condition).Find(&result)
		if err != nil{
			return nil,err
		} else {
			return result,err
		}
	} else {
		return nil,errors.New("invalid condition")
	}
}

// Entry DB's Dao

func CreateEntry(db *gorm.DB, entry *Entry,res interface{}) (int64,error){
	// Title Only
	result := db.Create(entry)

	FindEntry(db,"title = ?",entry.Title,1,res)

	return result.RowsAffected,result.Error
}

func FindEntry(db *gorm.DB, condition string,value string,number int,res interface{})  *gorm.DB{
	result := search(db.Model(&Entry{}),condition,value,number,res)
	return result
}

func UpdateEntry(db *gorm.DB,m *Entry,userid uint){
	err := db.Transaction(func(tx *gorm.DB) error {
		_, err := CreateHistory(tx, m, userid,&History{})
		if err != nil {
			return err
		}
		res := tx.Model(m).Updates(m)
		return res.Error
	})
	if err != nil {
		return 
	}
}
func DeleteEntry(db *gorm.DB,condition string,value string,number int)  (int64,error){
	res := FindEntry(db,condition,value,number,&Entry{})
	if res.Error == nil{
		result := res.Delete(&Entry{})
		return result.RowsAffected,nil
	} else {
		return 0,res.Error
	}
}

// History

func CreateHistory(db *gorm.DB,e *Entry,userid uint,res interface{}) (int64,error){
	result := db.Create(&History{
		EntryID: e.ID,
		UserID:  userid,
		Content: e.Content,
	})

	FindHistory(db,"Content = ?",e.Content,1,res)
	return result.RowsAffected,result.Error
}

func FindHistory(db *gorm.DB,condition string,value string,number int,res interface{}) *gorm.DB{
	result := search(db.Model(&History{}),condition,value,number,res)

	return result
}

func DropHistory(db *gorm.DB,condition string,value string,number int) (int64,error){

	result := FindHistory(db,condition,value,number,&History{})

	if result.Error != nil{
		return 0,result.Error
	} else {
		result := result.Delete(&History{})
		return result.RowsAffected,result.Error
	}
}

// Cat

func CreateCat(db *gorm.DB,p string,r interface{}) *gorm.DB {

	res := db.FirstOrCreate(&Cat{Path: p})
	SearchCat(db,"path = ?",p,r)
	return res
}

func SearchCat(db *gorm.DB,condition string,value string,res interface{}) *gorm.DB{
	result := search(db.Model(&Cat{}),condition,value,2,res)
	return result
}

func CatNode(db *gorm.DB,catNode string)  (Cat,error){
	var res Cat
	result := db.Where("path ~ ?",catNode).First(&res)
	return res,result.Error
}

func CatChildren(db *gorm.DB,catNode string,res interface{}) *gorm.DB{
	result := SearchCat(db,"path ~ ?",catNode+".*{1}",res)
	return result
}

func FindCat(db *gorm.DB,catNode string,res interface{}) *gorm.DB{
	result := SearchCat(db,"path <@ ?",catNode,res)
	return result
}

func CatParent(db *gorm.DB,catNode string,res interface{})  *gorm.DB {
	var result []Cat
	h := SearchCat(db,"path ~ ?","*{1}."+catNode,&result)

	if h.Error != nil{
		return h
	}

	comma := strings.Index(result[0].Path,".")
	p := result[0].Path[0:comma]
	h = SearchCat(db,"path ~ ?",p,res)
	return h
}

func CatBrother(db *gorm.DB,catNode string,res interface{}) *gorm.DB{
	var p []Cat
	h := CatParent(db,catNode,&p)
	if h.Error != nil{
		return h
	}
		parent := p[0].Path
		h = SearchCat(db,"path ~ ?",parent+"."+"!"+catNode+"{1}",&res)
		return h
}

func Cat2Entries(db *gorm.DB,c *Cat) ([]Entry,error){
	var res []Entry
	r := db.Model(c).Association("Entries").Find(&res)
	if r == nil{
		return res,r
	} else {
		return nil,r
	}
}

func DeleteCat(db *gorm.DB,c *Cat) (int64,error){
	res := db.Delete(c)
	return res.RowsAffected,res.Error
}

// Tag

func CreateTag(db *gorm.DB,t *Tag) (int64,error){
	res := db.Where("name = ?",t.Name).First(t)
	return res.RowsAffected,res.Error
}

func DeleteTag(db *gorm.DB,t *Tag) (int64,error){
	res := db.Delete(t)
	return res.RowsAffected,res.Error
}

// Group

func CreateGroup(db *gorm.DB,group []UserGroup) (int,error){
	for index,value := range group{
		result := db.Where(UserGroup{GroupName: value.GroupName}).FirstOrCreate(&value)
		if result.Error != nil{
			return index,result.Error
		}
	}
	return len(group),nil
}

func DeleteGroup(db *gorm.DB, id uint) (int64,error){
	if result := db.Delete(&UserGroup{
		Model:gorm.Model{
			ID: id,
		},
	}) ;result.Error != nil{
		return 0,result.Error
	} else {
		return result.RowsAffected,nil
	}
}