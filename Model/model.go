package Model

import (
	"gorm.io/gorm"
	"time"
)

type Entry struct {
	gorm.Model
	Title   string
	UserID  string
	Tags    string
	Cat     string
	History []History
	Content string
}

type History struct {
	gorm.Model
	EntryID uint
	UserID  string
	Content string
}

type UserGroup struct {
	gorm.Model
	GroupName string
	Users     []User
	Level     int
}
type User struct {
	CreateAt time.Time
	UpdateAt time.Time
	gorm.DeletedAt
	ID          string `gorm:"primaryKey"`
	Name        string
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

type RealName struct {
	id   uint
	Name string
}

//type Comments struct {
//	gorm.Model
//	Content string
//	UserID  uint
//}
