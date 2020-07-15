package main

import (
	"fmt"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	 "github.com/mrbardia72/blog-gorm-gorilla/src/routers"
	 "github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
) 

func main() {

	fmt.Println("start run server for API .......")
	helpers.InitialMigration()
	routers.HandleRequests()
}