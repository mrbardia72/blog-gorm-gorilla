package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"../model"
	"../config"
	"../helpers"
)

func Alluser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	var users []model.User

	db.Preload("Posts").Find(&users) //once user multi posts 1:n
	info:="get all users"
	helpers.GetTimeDate(info)

	json.NewEncoder(w).Encode(users)
}

func Newuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	info:=" New User Successfully Created "
	helpers.GetTimeDate(info)

	db.Create(&model.User{Name: name, Email: email})
	json.NewEncoder(w).Encode(&model.User{Name: name, Email: email})
}

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
