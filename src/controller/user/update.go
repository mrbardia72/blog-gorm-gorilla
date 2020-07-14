package user

import (
	"fmt"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Updateuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user model.User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	info:=" Successfully Updated User "
	helpers.LogApi(info)

	fmt.Fprintf(w, "Successfully Updated User")
}