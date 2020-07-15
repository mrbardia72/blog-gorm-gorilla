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

func Updatepost(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	title := vars["title"]
	body := vars["body"]

	var post model.Post
	err := db.Where("title = ?", title).Find(&post).Error

	if err != nil {
		
		errorx:=helpers.Errorx{Msgx:"not exisits record for update",Codex:"404"}
		json.NewEncoder(w).Encode(errorx)
		info:=" not exisits record for update "
		helpers.LogApi(info)
		
	} else {

		post.Body = body
		db.Save(&post)
		errorx:=helpers.Errorx{Msgx:"Successfully update post",Codex:"200"}
		json.NewEncoder(w).Encode(errorx)
		info:=" Successfully update post "
		helpers.LogApi(info)
	}
}