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

func Searchuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	name := vars["name"]

	var user model.User
	err := db.Where("name = ?", name).Find(&user).Error

	if err != nil {
		
		errorx:=helpers.Errorx{Msgx:"There is no such record",Codex:"404"}
		json.NewEncoder(w).Encode(errorx)
		info:="There is no such record"
		helpers.LogApi(info)
		
	} else { 
		
		db.Find(&user)
		json.NewEncoder(w).Encode(&user)
		info:=" Successfully update User "
		helpers.LogApi(info)
	}
}