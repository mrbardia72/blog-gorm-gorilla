package controller

import (
	"fmt"
	"net/http"
	// "encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
 
// func Newuser(w http.ResponseWriter, r *http.Request) {

// 	db := config.MySql()

// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	email := vars["email"]

// 	info:=" New User Successfully Created "
// 	helpers.GetTimeDate(info)

// 	db.Create(&model.User{Name: name, Email: email})
// 	json.NewEncoder(w).Encode(&model.User{Name: name, Email: email})
// }

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
	helpers.GetTimeDate(info)

	fmt.Fprintf(w, "Successfully Updated User")
}
