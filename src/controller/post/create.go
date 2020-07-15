package post

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
) 

func Newpost(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()
 
	vars := mux.Vars(r)
	title := vars["title"]
	body := vars["body"]

	info:=" New post Successfully Created "
	helpers.LogApi(info)

	db.Create(&model.Post{Title: title, Body: body})
	json.NewEncoder(w).Encode(&model.Post{Title: title, Body: body})
}