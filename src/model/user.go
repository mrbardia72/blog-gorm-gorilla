package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
	Posts []Post `gorm:"ForeignKey:UserID"`
}
