package post

import (
	"net/http"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Allpost(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	var posts []model.Post
	db.Preload("Posts").Find(&posts)
	
	info:="get all posts"
	helpers.LogApi(info)

	json.NewEncoder(w).Encode(posts)
}