package Model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Entry struct {
	gorm.Model
	Title string
	Author []User `gorm:"many2many:user_entries;"`
	Content string
}

type User struct {
	gorm.Model
	Name string
	Email string
	Pwd string
	Site string
	Country string
	Language string
	Entries	[]Entry `gorm:"many2many:user_entries;"`
	Comments []Comments `gorm:"ForeignKey:UserID;"`
	Mechanism string
	Sex bool
}

type Comments struct {
	gorm.Model
	Content string
	UserID uint
}