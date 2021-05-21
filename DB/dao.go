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

func search (db *gorm.DB,query string,value string,number int) ([]map[string]interface{},*gorm.DB,error){
	var r *gorm.DB
	r = db.Where(query,value)
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

// User DB's Dao

func RegisterUser(db *gorm.DB,user *User) (int64,error) {
	// Email only
	res := db.Create(user)
	return res.RowsAffected,res.Error
}

func UpdateUser(db *gorm.DB,query string,value string,number int,user *User) (int64,error) {
	_,h,err := FindUser(db,query,value,number)
	if err == nil{
		if res := h.Updates(user); res.Error != nil{
			return 0,res.Error
		} else {
			return res.RowsAffected,nil
		}
	} else {
		return 0,nil
	}
}

func DeleteUser(db *gorm.DB, condition string, value string, number int) (int64,error) {
	ans,r,err := FindUser(db,condition,value,number)
	if ans == nil {
		return 0, errors.New("No valid user")
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

func CreateEntry(db *gorm.DB, entry *Entry) (int64,error){
	// Title Only
	res := db.Where(Entry{Title: entry.Title}).Session(&gorm.Session{FullSaveAssociations: true}).FirstOrCreate(entry)
	return res.RowsAffected,res.Error
}

func FindEntry(db *gorm.DB, condition string,value string,number int)  ([]Entry,*gorm.DB,error){
	var result []Entry
	if r := db.Where(condition,value);r.Error == nil{
		singleEntry := Entry{}
		var lotsOfEntry []Entry
		switch number {
		case 1:
			r.First(&singleEntry)
			db.Model(&singleEntry).Preload("History").Find(&result)
		case 0:
			r.Find(&lotsOfEntry)
			db.Model(&lotsOfEntry).Preload("History").Find(&result)
		default:
			return nil,r,errors.New("wrong")
		}
		return result,r,nil
	} else {
		return nil,r,r.Error
	}
}

func UpdateEntry(db *gorm.DB,m *Entry,userid uint){
	db.Transaction(func(tx *gorm.DB) error{
		_, err := CreateHistory(tx,m,userid)
		if err != nil {
			return err
		}
		res := tx.Model(m).Updates(m)
		return res.Error
	})
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

// History

func CreateHistory(db *gorm.DB,e *Entry,userid uint) (int64,error){
	res := db.Create(&History{
		EntryID: e.ID,
		UserID:  userid,
		Content: e.Content,
	})

	return res.RowsAffected,res.Error
}

func DropHistory(db *gorm.DB,condition string,value string,number int) (int64,error){
	m := db.Model(&History{})
	_,h,err := search(m,condition,value,number)

	if err != nil{
		return 0,err
	} else {
		result := h.Delete(&History{})
		return result.RowsAffected,result.Error
	}
}

// Cat

func CreateCat(db *gorm.DB,p string) (Cat,*gorm.DB) {
	r := Cat{}
	res := db.Where(Cat{
		Path:    p,
	}).FirstOrCreate(&r)
	return r,res
}

func SearchCat(db *gorm.DB,condition string,value string) ([]Cat,error){
	var res []Cat
	result := db.Where(condition,value).Find(&res)
	return res,result.Error
}

func CatNode(db *gorm.DB,catNode string)  (Cat,error){
	var res Cat
	result := db.Where("path ~ ?",catNode).First(&res)
	return res,result.Error
}

func CatChildren(db *gorm.DB,catNode string) ([]Cat,error){
	res,err := SearchCat(db,"path ~ ?",catNode+".*{1}")
	return res,err
}

func CatFind(db *gorm.DB,catNode string) ([]Cat,error){
	res,err := SearchCat(db,"path <@ ?",catNode)
	return res, err
}

func CatParent(db *gorm.DB,catNode string) ([]Cat,error)  {
	res,err := SearchCat(db,"path ~ ?","*{1}."+catNode)
	comma := strings.Index(res[0].Path,".")
	p := res[0].Path[0:comma]
	res,err = SearchCat(db,"path ~ ?",p)
	return res,err
}

func CatBrother(db *gorm.DB,catNode string) ([]Cat,error){
	if p,err := CatParent(db,catNode);err == nil{
		parent := p[0].Path
		p,err = SearchCat(db,"path ~ ?",parent+"."+"!"+catNode+"{1}")

		return p,err
	} else {
		return p,err
	}
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