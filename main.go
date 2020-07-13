package main

import (
	"fmt"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	 "./src/model"
	 "./src/routers"
	 "./src/config"
)

func initialMigration() {

	db := config.MySql()
	db.AutoMigrate(&model.User{},&model.Post{})
}

func main() {

	fmt.Println("start API.......")
	initialMigration()
	routers.HandleRequests()
}