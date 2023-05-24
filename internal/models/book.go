package models

import(
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
	Description string `json:"description"`
	Uploader string `json:"uploader"`
}