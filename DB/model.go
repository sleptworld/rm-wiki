package DB

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	gorm.Model
	Name    string
	Entries []*Entry `gorm:"many2many:tag_entries"`
}

type Cat struct {
	ID          int32
	Path        ltree  `gorm:"type:ltree"`
	Entries     []Entry
}

type Entry struct {
	gorm.Model
	Title     string
	UserID    uint `gorm:"<-:create"`
	Lock      bool `gorm:"not null;default:false"`
	EditingBy uint
	Tags      []*Tag `gorm:"many2many:tag_entries"`
	CatID     uint
	History   []History
	Content   string `gorm:"size:15000"`
	Info      string `gorm:"default:Create;size:30"`
}

type History struct {
	gorm.Model
	EntryID uint   `gorm:"<-:create"`
	UserID  uint   `gorm:"<-:create"`
	Content string `gorm:"size:15000"`
	Info    string `gorm:"default:Updated;size:30"`
}

type Draft struct {
	gorm.Model
	UserID  uint   `gorm:"<-:create"`
	Content string `gorm:"size:15000"`
}

type UserGroup struct {
	gorm.Model
	GroupName string
	Users     []User
	Level     int8 `gorm:"not null"`
}
type User struct {
	gorm.Model
	Name        string `gorm:"not null;check:Name <> '';size:30"`
	Email       string `gorm:"unique;check:Email <> '';size:30"`
	Pwd         []byte `gorm:"not null;type:bytea"`
	UserGroupID uint   `gorm:"not null;default:2"`
	Birthday    time.Time
	Avatar      string
	Description string `gorm:"default:'He said ohhhhhh'"`
	Site        string `gorm:"size:30"`
	Country     string `gorm:"size:20"`
	Language    string `gorm:"size:5"`
	Entries     []Entry
	EditEntries []History
	Drafts      []Draft
	Mechanism   string `gorm:"size:30"`
	Sex         int8   `gorm:"check:Sex IN (0,1,2)"`
	Profession  string `gorm:"size:30"`
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
