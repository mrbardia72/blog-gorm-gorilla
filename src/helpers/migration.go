package helpers

import (
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	 "github.com/mrbardia72/blog-gorm-gorilla/src/model"
	 "github.com/mrbardia72/blog-gorm-gorilla/src/config"
) 

func InitialMigration() {

	db := config.MySql()
	db.AutoMigrate(&model.User{},&model.Post{})
}