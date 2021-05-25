package DB

import (
	"database/sql/driver"
)

type Level = int8

type ltree struct {
	path string
}

func (tree *ltree) Scan(value interface{}) error{
	tree.path = value.(string)
	return nil
}

func (tree ltree) Value() (driver.Value, error){
	if tree.path == ""{
		return nil,nil
	}
	return tree.path,nil
}