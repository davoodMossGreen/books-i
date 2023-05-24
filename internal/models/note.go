package models

import(
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Heading		string
	Note		string
	Username	string
}