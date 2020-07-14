package controller


import (
	// "fmt"
	"net/http"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Allpost(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	var post []model.Post
	// db.Model(&user).Related(&post)
	// //// SELECT * FROM posts WHERE user_id = 111; // 111 is user's ID
	
	db.Find(&post)

	info:="get all posts"
	helpers.LogApi(info)

	json.NewEncoder(w).Encode(post)
}
