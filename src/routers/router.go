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
	s.HandleFunc("/users", user.Alluser).Methods("GET")
	s.HandleFunc("/user/{name}", user.Deleteuser).Methods("DELETE")
	s.HandleFunc("/user/{name}/{email}", user.Updateuser).Methods("PUT")
	s.HandleFunc("/user/{name}/{email}", user.Newuser).Methods("POST")

	//routers post
	s.HandleFunc("/posts", controller.Allpost).Methods("GET")	//v1/api/infopost/posts

	log.Fatal(http.ListenAndServe(":8081", r))
}