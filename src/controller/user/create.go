package user

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
) 

func Newuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()
 
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	info:=" New User Successfully Created "
	helpers.LogApi(info)

	db.Create(&model.User{Name: name, Email: email})
	json.NewEncoder(w).Encode(&model.User{Name: name, Email: email})
}