package post

import (
	"net/http"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
 
func Deletepost(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	title := vars["title"]

	var post model.Post
	err := db.Where("title = ?", title).Find(&post).Error

	if err != nil {
		
		errorx:=helpers.Errorx{Msgx:"not exisits record for delete",Codex:"404"}
		json.NewEncoder(w).Encode(errorx)
		info:=" not exisits record for delete "
		helpers.LogApi(info)
		
	} else {

		db.Delete(&post)
		errorx:=helpers.Errorx{Msgx:"Successfully Deleted post",Codex:"200"}
		json.NewEncoder(w).Encode(errorx)
		info:=" Successfully Deleted post "
		helpers.LogApi(info)
	}
}