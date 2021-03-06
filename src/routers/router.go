package routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/mrbardia72/blog-gorm-gorilla/src/controller/post"
	"github.com/mrbardia72/blog-gorm-gorilla/src/controller/user"
)

func HandleRequests() {

	route := mux.NewRouter()

	userroute := route.PathPrefix("/v1/api/user").Subrouter()
	userroute.HandleFunc("/read", user.Alluser).Methods("GET")
	userroute.HandleFunc("/delete/{name}", user.Deleteuser).Methods("DELETE")
	userroute.HandleFunc("/update/{name}/{email}", user.Updateuser).Methods("PUT")
	userroute.HandleFunc("/create/{name}/{email}", user.Newuser).Methods("POST")
	userroute.HandleFunc("/search/{name}", user.Searchuser).Methods("GET")

	postroute := route.PathPrefix("/v1/api/post").Subrouter()
	postroute.HandleFunc("/read", post.Allpost).Methods("GET")
	postroute.HandleFunc("/delete/{title}", post.Deletepost).Methods("DELETE")
	postroute.HandleFunc("/update/{title}/{body}", post.Updatepost).Methods("PUT")
	postroute.HandleFunc("/create/{title}/{body}", post.Newpost).Methods("POST")
	postroute.HandleFunc("/search/{title}", post.Searchpost).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", route))
}