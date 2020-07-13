package config

import (
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
)

func MySql()(*gorm.DB)  {

	dbDriver:="mysql"
	dbName:="digi"
	dbUser:="root"
	dbPassword:="Bardi@1993"

	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}