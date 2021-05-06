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

type User struct {
	gorm.Model
	Name        string
	Email       string
	Pwd         string
	Avatar      string `gorm:"default:empty.jpg"`
	Description string `gorm:"default:Ahhhh"`
	Site        string `gorm:"default:empty"`
	Country     string `gorm:"default:China"`
	Language    string `gorm:"default:Chinese(Simpled)"`
	Entries     []Entry
	EditEntries []History
	Mechanism   string `gorm:"default:empty"`
	Sex         bool   `gorm:"default:0"`
	Profession  string `gorm:"default:empty"`
}

//type Comments struct {
//	gorm.Model
//	Content string
//	UserID  uint
//}
