package main

import (
	"fmt"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	//  "github.com/mrbardia72/blog-gorm-gorilla/src/model"
	 "github.com/mrbardia72/blog-gorm-gorilla/src/routers"
	//  "github.com/mrbardia72/blog-gorm-gorilla/src/config"
	 "github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
) 

// func initialMigration() {

// 	db := config.MySql()
// 	db.AutoMigrate(&model.User{},&model.Post{})
// }

func main() {

	fmt.Println("start run server for API .......")
	helpers.InitialMigration()
	routers.HandleRequests()
}