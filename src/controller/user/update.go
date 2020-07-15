package user

import (
	"encoding/json"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

func Updateuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user model.User
	err := db.Where("name = ?", name).Find(&user).Error

	if err != nil {
		
		errorx:=helpers.Errorx{Msgx:"not exisits record for update",Codex:"404"}
		json.NewEncoder(w).Encode(errorx)
		info:=" not exisits record for update "
		helpers.LogApi(info)
		
	} else {

		user.Email = email
		db.Save(&user)
		errorx:=helpers.Errorx{Msgx:"Successfully update User",Codex:"200"}
		json.NewEncoder(w).Encode(errorx)
		info:=" Successfully update User "
		helpers.LogApi(info)
	}
}