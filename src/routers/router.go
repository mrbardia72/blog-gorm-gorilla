package routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/mrbardia72/blog-gorm-gorilla/src/controller"
	"github.com/mrbardia72/blog-gorm-gorilla/src/controller/user"
)

func HandleRequests() {

	route := mux.NewRouter()

	userroute := route.PathPrefix("/v1/api/u.er").Subrouter()
	userroute.HandleFunc("/read", user.Alluser).Methods("GET")
	userroute.HandleFunc("/delete/{name}", user.Deleteuser).Methods("DELETE")
	userroute.HandleFunc("/update/{name}/{email}", user.Updateuser).Methods("PUT")
	userroute.HandleFunc("/create/{name}/{email}", user.Newuser).Methods("POST")
	userroute.HandleFunc("/search/{name}", user.Searchuser).Methods("GET")

	postroute := route.PathPrefix("/v1/api/u.er").Subrouter()
	postroute.HandleFunc("/posts", controller.Allpost).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", route))
}