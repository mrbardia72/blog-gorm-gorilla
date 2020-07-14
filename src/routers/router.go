package routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/mrbardia72/blog-gorm-gorilla/src/controller"
	"github.com/mrbardia72/blog-gorm-gorilla/src/controller/user"
)
func HandleRequests() {

	r := mux.NewRouter()
	s := r.PathPrefix("/v1/api").Subrouter()

	//routers user
	s.HandleFunc("/users", user.Alluser).Methods("GET")	//v1/api/infouser/users
	r.HandleFunc("/user/{name}", controller.Deleteuser).Methods("DELETE")
	r.HandleFunc("/user/{name}/{email}", controller.Updateuser).Methods("PUT")
	r.HandleFunc("/user/{name}/{email}", controller.Newuser).Methods("POST")

	//routers post
	s.HandleFunc("/posts", controller.Allpost).Methods("GET")	//v1/api/infopost/posts

	log.Fatal(http.ListenAndServe(":8081", r))
}