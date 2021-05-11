package Model

import (
	"fmt"
	"github.com/goinggo/mapstructure"
)

func IsContain(items []string , item string) bool{
	for _,eachitem := range items{
		if eachitem == item{
			return true
		}
	}
	return false
}
func MapToStruct(m map[string]interface{},s interface{}) (error){
	if err := mapstructure.Decode(m,s);err != nil{
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}