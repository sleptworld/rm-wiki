package DB

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	gorm.Model
	Name    string  `gorm:"size:20;unique"`
	Entries []Entry `gorm:"many2many:tag_entries"`
}

type Cat struct {
	ID      int32   `gorm:"primaryKey;autoIncrement"`
	Path    string  `gorm:"type:ltree;unique"`
	Entries []Entry `gorm:"foreignKey:CatID"`
}

type Lang struct {
	ID      int8    `gorm:"primaryKey;autoIncrement"`
	Lang    string  `gorm:"size:5;uniqueIndex"`
	Entries []Entry `gorm:"foreignKey:LangID"`
}

type Entry struct {
	// gorm Model
	gorm.Model
	//
	Title   string `gorm:"size:30;<-:create;not null;check:title <> '';unique"`
	Content string `gorm:"size:15000"`
	Info    string `gorm:"default:Create;size:30"`
	Review  bool   `gorm:"default:false;not null"`
	// one2many
	Tags    []Tag `gorm:"many2many:tag_entries"`
	History []History
	// ForeignKey
	LangID int8  `gorm:"default:1"`
	UserID uint  `gorm:"<-:create"`
	CatID  int32 `gorm:"not null;default:1"`
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
	UserID  uint `gorm:"<-:create"`
	Title   string
	Content string `gorm:"size:15000"`
}

type UserGroup struct {
	gorm.Model
	GroupName string `gorm:"size:10"`
	Users     []User
	Level     Level  `gorm:"not null;default:0"`
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