package DB

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/lib/pq"
	"strconv"
)

var (
	a *gormadapter.Adapter
	E *casbin.Enforcer
	err error
)

func InitCasbin(){
	a, err = gormadapter.NewAdapterByDB(Db)
	if err != nil{
		panic(err)
	}
	E, err = casbin.NewEnforcer("model.conf", a)
	if err != nil{
		panic(err)
	}
	err := E.LoadPolicy()
	if err != nil {
		return 
	}
}

func Check(sub uint,obj string,act string) bool {
	enforce, err := E.Enforce(strconv.Itoa(int(sub)), obj, act)
	if err != nil {
		return false
	}
	return enforce
}

func AddRoleForUser(UserID uint,Group string) bool {
	user, err := E.AddRoleForUser(strconv.Itoa(int(UserID)), Group)
	if err != nil {
		return false
	} else {
		return user
	}
}

func DeleteRoleForUser(UserID uint,Group string) bool{
	user,err := E.DeleteRoleForUser(strconv.Itoa(int(UserID)),Group)
	if err != nil{
		return false
	} else {
		return user
	}
}

func GetRolesByUser(UserID uint) []string  {
	user, _ := E.GetRolesForUser(strconv.Itoa(int(UserID)))
	return user
}

func GetAllRoles() []string {
	roles := E.GetAllRoles()
	return roles
}

func GetAllUserByRole(Group string) []string  {
	res,_ := E.GetUsersForRole(Group)
	return res
}

func IsThisRole(UserID uint,Group string) bool {
	user, err := E.HasRoleForUser(strconv.Itoa(int(UserID)), Group)
	if err != nil {
		return false
	} else {
		return user
	}
}

func AddRole(Group [][]string) bool  {
	policies, err := E.AddPolicies(Group)
	if err != nil {
		return false
	} else {
		return policies
	}
}