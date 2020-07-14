package user

import (
	"net/http"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Alluser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	var users []model.User

	db.Preload("Posts").Find(&users) //once user multi posts 1:n
	
	info:="get all users"
	helpers.GetTimeDate(info)

	json.NewEncoder(w).Encode(users)
}