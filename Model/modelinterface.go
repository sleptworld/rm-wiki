package Model
//
//import (
//	"errors"
//	"github.com/sleptworld/test/DB"
//	"github.com/sleptworld/test/tools"
//	"gorm.io/gorm"
//)
//
//type Model struct {
//	to interface{}
//}
//
//func (m *Model) reflectConvert(from interface{},to interface{}) (interface{},error){
//	res,err := tools.ReadLabel(from,to)
//	return res,err
//}
//
//func (m *Model) Init(from interface{},to interface{}){
//	if c,err := m.reflectConvert(from,to);err == nil {
//		m.to = c
//	}
//}
//
//func (m *Model) Get() interface{} {
//	return m.to
//}
//
//
//type UserModel struct {
//	Model
//}
//
//func (u *UserModel) Create(res interface{}) (int64,error){
//	if t,ok := u.Get().(DB.Entry);ok{
//		num, err2 := t.Create(res:
//		if err2 != nil{
//			return 0,err2
//		}
//		return num,nil
//	}
//	return 0,errors.New("Assert")
//}
//
//func (u *UserModel) Update(re interface{}) (int64,error){
//	if t,ok := u.Get().(DB.Entry);ok{
//		num, err2 := t.Update(re)
//		if err2 != nil {
//			return 0,err2
//		}
//		return num,nil
//	}
//	return 0,errors.New("")
//}
//
//func (u *UserModel) Delete(c string,values ...interface{}) (int64,error){
//	if t,ok:=u.Get().(DB.Entry);ok{
//		i, err := t.Delete(nil)
//		if err != nil {
//			return
//		}
//		return
//	}
//	return
//}
//
//func (u *UserModel) Query(re interface{},number int,c string,values ...interface{}) *gorm.DB{
//	if t,ok := u.Get().(DB.Entry);ok{
//		res := t.Query(re,number,c,values...)
//		return res
//	}
//	return nil
//}
