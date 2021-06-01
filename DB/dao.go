package DB

import (
	"errors"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"reflect"
	"strings"
)


type Model struct {
	m interface{}
}

func Init(s interface{}) *Model{
	reType := reflect.TypeOf(s)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct || reType.Elem().Kind() != reflect.Slice{
		return nil
	}

	return &Model{m: s}
}

func (m *Model) Init (s interface{}) *Model{

	m.m = s
	return m
}

func (m *Model) Query(re interface{},number int,c string,values ...interface{})*gorm.DB{
	if c == ""{
		return search(Db.Model(m.m), number,re,m.m,nil)
	}
	return search(Db.Model(m.m),number,re,c,values)
}

func (m *Model) Create(re interface{}) (int64,error){
	res := Db.Where(m.m).FirstOrCreate(m.m)
	m.Query(re,1,"",nil)
	return res.RowsAffected,res.Error
}

func (m *Model) Update(re interface{}) (int64,error){
	res := Db.Model(m.m).Updates(m.m)
	return res.RowsAffected,res.Error
}

func (m *Model) Delete(c string,values ...interface{}) (int64,error){
	res := delete(m.m,c,values...)
	return res.RowsAffected,res.Error
}

func search (db *gorm.DB,number int , res interface{},query interface{},value ...interface{}) *gorm.DB{
	var r *gorm.DB
	if reflect.TypeOf(query).Kind() != reflect.String{
		r = db.Where(query)
	}else {
		r = db.Where(query,value...)
	}
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

// tools for delete

func delete(m interface{},c string,values ...interface{})  *gorm.DB{
	var result *gorm.DB
	if c == ""{
		result = Db.Where(m).Delete(m)
	} else {
		result = Db.Where(c,values...).Delete(m)
	}
	return result
}

// Cat

func (cat *Cat) Node(re interface{})  *gorm.DB{
	result := Db.Where("path ~ ?",cat.Path).First(&re)
	return result
}

func (cat *Cat) Children(res interface{}) *gorm.DB{

	result := Init(cat).Query(res,3,"path ~ ?",cat.Path+".*{1}",res)
	return result
}

func (cat *Cat) Find (res interface{}) *gorm.DB{
	result := Init(cat).Query(res,2,"path <@ ?",cat.Path,res)
	return result
}

func (cat *Cat) Parent(res interface{})  *gorm.DB {
	var result []Cat

	h := Init(cat).Query(&result,2,"path ~ ?","*{1}."+cat.Path)
	if h.Error != nil{
		return h
	}

	comma := strings.Index(result[0].Path,".")
	p := result[0].Path[0:comma]
	h = Init(cat).Query(res,2,"path ~ ?",p)
	return h
}

func (cat *Cat) Brother(res interface{}) *gorm.DB{
	var p []Cat
	h := cat.Parent(&p)
	if h.Error != nil{
		return h
	}
	parent := p[0].Path
	h = Init(cat).Query(res,2,"path ~ ?",parent+"."+"!"+cat.Path+"{1}")
	return h
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


func Cat2Entries(db *gorm.DB,c *Cat) ([]Entry,error){
	var res []Entry
	r := db.Model(c).Association("Entries").Find(&res)
	if r == nil{
		return res,r
	} else {
		return nil,r
	}
}