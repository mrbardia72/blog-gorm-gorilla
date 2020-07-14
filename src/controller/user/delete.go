package user

import (

	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
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
	err := db.Where("name = ?", name).Find(&user).Error

	if gorm.IsRecordNotFoundError(err) {
		fmt.Fprintf(w, "not exisits record for delete")

	} else {
	db.Delete(&user)
	}

	info:=" Successfully Deleted User "
	helpers.GetTimeDate(info)

	fmt.Fprintf(w, "Successfully Deleted User")
}