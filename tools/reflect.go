package tools

import (
	"errors"
	"reflect"
)

const (
	label = "DBModel"
)

func ReadLabel(s interface{},ts interface{}) (interface{},[]string,error)  {

	value := reflect.ValueOf(s)
	reType := reflect.TypeOf(ts)

	var errField []string

	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct ||
		reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct{

		return nil,errField,errors.New("Paras should be the ptr of Struct.")

	}

	to := reflect.New(reType.Elem()).Elem()
	value2 := value.Elem()

	for i:=0;i<value2.Type().NumField();i++{
		var result string

		structField := value2.Type().Field(i)
		name := structField.Name
		tag := structField.Tag
		if t := tag.Get("DBModel"); t == ""{
			result = name
		} else {
			result = t
		}
		if tmp := to.FieldByName(result);tmp.CanSet() && tmp.Kind() == structField.Type.Kind(){
			tmp.Set(value2.Field(i))
		} else {
			errField = append(errField,result)
		}
	}

	return to.Interface(),errField,nil
}

func ReflectGetValue(s interface{},key string) interface{} {
	values := reflect.ValueOf(s)

	if values.Kind() != reflect.Ptr || values.Elem().Kind() != reflect.Struct{
		return nil
	}
	return values.Elem().FieldByName(key).Interface()

}