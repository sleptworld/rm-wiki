package Controller

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	a gormadapter.Adapter
)

func Init(){
	a , _ = gormadapter.NewAdapter("postgres","")
	e , _ := casbin.NewEnforcer("/path/.conf",a)

}