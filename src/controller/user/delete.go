package user

import (

	"fmt"
	"net/http"
	"encoding/json"
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

	if err != nil {

		fmt.Fprintf(w, "not exisits record for delete")
		info:=" not exisits record for delete "
		helpers.GetTimeDate(info)
		
	} else {

		db.Delete(&user)
		info:=" Successfully Deleted User "
		json.NewEncoder(w).Encode(info)
		helpers.GetTimeDate(info)
		fmt.Fprintf(w, "Successfully Deleted User")
	}
}