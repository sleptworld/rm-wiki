package Model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func RegisterUser(db *gorm.DB,user *User)  {
	result := db.Create(user)
	if result.Error != nil {
		fmt.Printf("Error")
		return
	}
}

func UpdateUser(db *gorm.DB,user *User)  {
	result := db.Updates(user)
	if result.Error != nil{
		fmt.Printf("Error")
		return
	}
}

func LogoutUser(db *gorm.DB, id uint)  {
	result := db.Delete(&User{},id)
	if result.Error == nil{
		allEntry := db.Where("UserID <> ?",id).Find(&Entry{})
		allHistory := db.Where("UserID <> ?",id).Find(&History{})
		if allHistory.Error == nil && allEntry.Error == nil{

		}

	}
}

func FindUser(db *gorm.DB,)  {

}