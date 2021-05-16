package DB

import (
	"gorm.io/gorm"
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
	Pwd         []byte
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

type RealName struct {
	id   uint
	Name string
}

//type Comments struct {
//	gorm.DB
//	Content string
//	UserID  uint
//}
