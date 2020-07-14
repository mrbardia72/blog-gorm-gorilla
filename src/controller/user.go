package controller

import (
	"fmt"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Deleteuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	name := vars["name"]

	var user model.User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	info:=" Successfully Deleted User "
	helpers.GetTimeDate(info)

	fmt.Fprintf(w, "Successfully Deleted User")
}
