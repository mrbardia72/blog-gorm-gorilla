package post

import (
	"encoding/json"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Searchpost(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	title := vars["title"]

	var post model.Post
	err := db.Where("title = ?", title).Find(&post).Error

	if err != nil {
		
		errorx:=helpers.Errorx{Msgx:"There is no such record",Codex:"404"}
		json.NewEncoder(w).Encode(errorx)
		info:="There is no such record"
		helpers.LogApi(info)
		
	} else { 
		
		db.Find(&post)
		json.NewEncoder(w).Encode(&post)
		info:=" Successfully update post "
		helpers.LogApi(info)
	}
}