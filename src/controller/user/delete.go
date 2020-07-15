package user

import (
	"net/http"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/mrbardia72/blog-gorm-gorilla/src/model"
	"github.com/mrbardia72/blog-gorm-gorilla/src/config"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
 
func Deleteuser(w http.ResponseWriter, r *http.Request) {

	db := config.MySql()

	vars := mux.Vars(r)
	name := vars["name"]

	var user model.User
	err := db.Where("name = ?", name).Find(&user).Error

	if err != nil {
		
		errorx:=helpers.Errorx{Msgx:"not exisits record for delete",Codex:"404"}
		json.NewEncoder(w).Encode(errorx)
		info:=" not exisits record for delete "
		helpers.LogApi(info)
		
	} else {

		db.Delete(&user)
		errorx:=helpers.Errorx{Msgx:"Successfully Deleted User",Codex:"200"}
		json.NewEncoder(w).Encode(errorx)
		info:=" Successfully Deleted User "
		helpers.LogApi(info)
	}
}