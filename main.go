package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	manageDatabaseSchema()
	initApi()
}

func manageDatabaseSchema() {
	db, err := getDataBase()
	if err != nil {
		log.Fatal("Couldn't connect to database ", err)
	}
	defer db.Close()
	if !db.HasTable(User{}) {
		createDefaultSchema(db)
	}

}

func initApi() {

	router := mux.NewRouter()
	sr := router.PathPrefix("/auth").Methods("POST").Subrouter()
	sr.HandleFunc("/signup", signup)
	sr.HandleFunc("/login", login)
	sr.HandleFunc("/change/passwd", changePasswd)
	log.Print("Listening at 8000....")
	http.ListenAndServe(":8000", router)
}
