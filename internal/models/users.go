package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username		string 		`json:"username" gorm:"unique"`
	Email 			string 		`json:"email" gorm:"unique"`
	Password 		string		`json:"password"`
	User_type 		string 		`json:"user_type"`
}
