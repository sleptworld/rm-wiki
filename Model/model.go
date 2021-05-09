package Model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Entry struct {
	gorm.Model
	Title   string
	UserID  uint
	Tags    string
	Cat     string
	History []History
	Content string
}

type History struct {
	gorm.Model
	EntryID uint
	UserID  uint
	Content string
}

type UserGroup struct {
	gorm.Model
	GroupName string
	Users     []User
	Level     int
}
type User struct {
	gorm.Model
	Name        string
	Email       string
	Pwd         string
	UserGroupID uint
	Avatar      string
	Description string
	Site        string
	Country     string
	Language    string
	Entries     []Entry
	EditEntries []History
	Mechanism   string
	Sex         int
	Profession  string
}

//type Comments struct {
//	gorm.Model
//	Content string
//	UserID  uint
//}
