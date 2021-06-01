package tools

import (
	"fmt"
	"reflect"
)

const (
	label = "DBModel"
)

func ReflectReadStruct (s interface{},ts interface{},errHandler func(value reflect.Value,key string,set *reflect.Value)) (interface{},[]string)  {

	value := reflect.ValueOf(s)
	reType := reflect.TypeOf(ts)

	var errField []string

	fmt.Println(value.Elem().Kind(),value.Kind(),reType.Kind(),reType.Elem().Kind())
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct ||
		reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct{

		return nil,errField

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
		if tmp := to.FieldByName(result);tmp.CanSet() && tmp.Type() == structField.Type{
			tmp.Set(value2.Field(i))
		} else {
			errHandler(value2.Field(i),result,&tmp)
			errField = append(errField,result)
		}
	}

	return to.Interface(),errField
}

func ReflectGetValue(s interface{},key string) interface{} {
	values := reflect.ValueOf(s)

	if values.Kind() != reflect.Ptr || values.Elem().Kind() != reflect.Struct{
		return nil
	}
	return values.Elem().FieldByName(key).Interface()
}

func ReflectSliceRead(s interface{},ts interface{},errHandler func(value reflect.Value,key string,set *reflect.Value)) ([]interface{},[][]string){
	value := reflect.ValueOf(s)
	reType := reflect.TypeOf(ts)
	var errField [][]string
	res := make([]interface{},0)

	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Slice || reType.Kind() != reflect.Ptr{
		return nil,errField
	}
	for i := 0;i<value.Elem().Len();i++{
		to := reflect.New(reType.Elem()).Elem()
		var errs []string
		t := value.Elem().Index(i)
		for i:=0;i<t.Type().NumField();i++{

			var result string
			structField := t.Type().Field(i)
			name := structField.Name
			tag := structField.Tag
			if t := tag.Get("DBModel"); t == ""{
				result = name
			} else {
				result = t
			}
			if tmp := to.FieldByName(result);tmp.CanSet() && tmp.Type() == structField.Type{
				tmp.Set(t.Field(i))
			} else {
				errs = append(errs,result)
				errHandler(t.Field(i),result,&tmp)
			}
		}
		errField = append(errField,errs)
		res = append(res,to.Interface())
	}
	return res,errField
}