package backup
//
//import (
//	"github.com/sleptworld/test/tools"
//	"gorm.io/gorm"
//	"strings"
//)
//
//// History
//
//func (h *History) Query(re interface{},number int,c string,values ...interface{}) *gorm.DB {
//	var res *gorm.DB
//	if c == ""{
//		res = search(Db.Model(&History{}),number,re,h,nil)
//	} else {
//		res = search(Db.Model(h),number,re,c,values...)
//	}
//	return res
//}
//
//func (h *History) create(e *Entry,re interface{}) (int64,error){
//
//	result := Db.Create(&History{
//		EntryID: e.ID,
//		UserID:  e.UserID,
//		Content: e.Content,
//	})
//
//	h.Query(re,1,"",nil)
//	return result.RowsAffected,result.Error
//}
//
//func (h *History) Delete(c string,values ...interface{}) (int64,error){
//	res := delete(h,c,values...)
//	return res.RowsAffected,res.Error
//}
//
//
//// Cat
//
//func (cat *Cat) Query(re interface{},number int,c string,values ...interface{}) *gorm.DB{
//	if c == ""{
//		res := search(Db.Model(cat),number,re,cat,nil)
//		return res
//	}
//	res := search(Db.Model(cat),number,re,c,values...)
//	return res
//}
//
//func (cat *Cat) Update(re interface{})  (int64,error){
//	return 0,nil
//}
//
//func (cat *Cat) Create(re interface{})(int64,error)  {
//	res := Db.Create(cat)
//	cat.Query(re,1,"",nil)
//	return  res.RowsAffected,res.Error
//}
//
//func (cat *Cat) Delete(c string,values ...interface{}) (int64,error)  {
//	res := delete(cat,c,values...)
//	return res.RowsAffected,res.Error
//}
//
//func (cat *Cat) Node(re interface{})  *gorm.DB{
//	result := Db.Where("path ~ ?",cat.Path).First(&re)
//	return result
//}
//
//func (cat *Cat) Children(res interface{}) *gorm.DB{
//	result := cat.Query(res,3,"path ~ ?",cat.Path+".*{1}",res)
//	return result
//}
//
//func (cat *Cat) Find (res interface{}) *gorm.DB{
//	result := cat.Query(res,2,"path <@ ?",cat.Path,res)
//	return result
//}
//
//func (cat *Cat) Parent(res interface{})  *gorm.DB {
//	var result []Cat
//
//	h := cat.Query(&result,2,"path ~ ?","*{1}."+cat.Path)
//	if h.Error != nil{
//		return h
//	}
//
//	comma := strings.Index(result[0].Path,".")
//	p := result[0].Path[0:comma]
//	h = cat.Query(res,2,"path ~ ?",p)
//	return h
//}
//
//func (cat *Cat) Brother(res interface{}) *gorm.DB{
//	var p []Cat
//	h := cat.Parent(&p)
//	if h.Error != nil{
//		return h
//	}
//	parent := p[0].Path
//	h = cat.Query(res,2,"path ~ ?",parent+"."+"!"+cat.Path+"{1}")
//	return h
//}
//
//// Entry
//
//func (e *Entry) Create(re interface{}) (int64,error){
//	result := Db.Create(e)
//	e.Query(re,1,"",nil)
//	return result.RowsAffected,result.Error
//}
//
//func (e *Entry) Delete(c string,values ...interface{}) (int64,error){
//	res := delete(e,c,values...)
//	return res.RowsAffected,res.Error
//}
//
//func (e *Entry) Update(re interface{}) (int64,error){
//	res := Db.Model(e).Updates(e)
//	e.Query(re,1,"",nil)
//	return res.RowsAffected,res.Error
//}
//
//func (e *Entry) Query(re interface{},number int,c string,values ...interface{}) *gorm.DB{
//	if c == ""{
//		result := search(Db.Model(e),number,re,e,nil)
//		return result
//	}
//	result := search(Db.Model(e),number,re,c,values...)
//	return result
//}
//
//
//
//// User DB's Dao
//
//func (u *User) Query(re interface{}, number int,c string,values ...interface{}) *gorm.DB{
//	if c == ""{
//		result := search(Db.Model(u),number,re,u,nil)
//		return result
//	}
//	result := search(Db.Model(u),number,re,c,values...)
//	return result
//}
//
//func (u *User) Create(re interface{}) (int64,error)  {
//	result := Db.Where(u).FirstOrCreate(u)
//	u.Query(re,1,"",nil)
//	return result.RowsAffected,result.Error
//}
//
//func (u *User) Delete(c string,values ...interface{}) (int64,error){
//	res := delete(u,c,values...)
//	return res.RowsAffected,res.Error
//}
//
//func (u *User) Update(re interface{}) (int64,error){
//	res := Db.Model(u).Updates(u)
//	u.Query(re,1,"",nil)
//	return res.RowsAffected,res.Error
//}
//
//func UserForeignKey(db *gorm.DB,m map[string]interface{},condition string) ([]map[string]interface{},error){
//	choice := []string{"Entries","EditEntries"}
//	userModel := User{
//		Model:gorm.Model{
//			ID: m["id"].(uint),
//		},
//	}
//	var result []map[string]interface{}
//	if tools.IsContain(choice,condition){
//		err := db.Model(&userModel).Association(condition).Find(&result)
//		if err != nil{
//			return nil,err
//		} else {
//			return result,err
//		}
//	} else {
//		return nil,errors.New("invalid condition")
//	}
//}
//
//
//func Cat2Entries(db *gorm.DB,c *Cat) ([]Entry,error){
//	var res []Entry
//	r := db.Model(c).Association("Entries").Find(&res)
//	if r == nil{
//		return res,r
//	} else {
//		return nil,r
//	}
//}
//
//// Tag
//
//func (tag *Tag) Create(re interface{}) (int64,error){
//	res := Db.Where(tag).FirstOrCreate(tag)
//	return res.RowsAffected,res.Error
//}
//
//
//func (tag *Tag) Delete(c string,values ...interface{}) (int64,error){
//	res := delete(tag,c,values...)
//	return res.RowsAffected,res.Error
//}
//
//func (tag *Tag) Update(re interface{}) (int64,error){
//	return 0,nil
//}
//
//func  (tag *Tag) Query(re interface{},number int,c string,values ...interface{}) *gorm.DB  {
//	if c == ""{
//		return search(Db.Model(tag),number,re,tag,nil)
//	}
//
//	return search(Db.Model(tag),number,re,c,values...)
//}
//
//// Group
//
//func (g *UserGroup) Create(re interface{}) (int64,error)  {
//	res := Db.Where(g).FirstOrCreate(g)
//	return res.RowsAffected,res.Error
//}
//
//func (g *UserGroup) Delete(c string,value ...interface{}) (int64,error){
//	res := delete(g,c,value...)
//	return res.RowsAffected,res.Error
//}
//
//func CreateGroup(db *gorm.DB,group []UserGroup) (int,error){
//	for index,value := range group{
//		result := db.Where(UserGroup{GroupName: value.GroupName}).FirstOrCreate(&value)
//		if result.Error != nil{
//			return index,result.Error
//		}
//	}
//	return len(group),nil
//}
//
//func DeleteGroup(db *gorm.DB, id uint) (int64,error){
//	if result := db.Delete(&UserGroup{
//		Model:gorm.Model{
//			ID: id,
//		},
//	}) ;result.Error != nil{
//		return 0,result.Error
//	} else {
//		return result.RowsAffected,nil
//	}
//}
