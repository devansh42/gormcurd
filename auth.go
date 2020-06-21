package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//signup, signups a new user
func signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
	}
	vals := r.Form

	hpasswd := vals.Get("password")
	email := vals.Get("email")
	username := vals.Get("username")
	log.Print(email, hpasswd, username)
	user := User{Email: email, UserName: username, Password: fmt.Sprint(hpasswd)} //New User
	db, err := getDataBase()
	if err != nil {
		//Handle Error
		w.WriteHeader(500)
		log.Print("Couldn't connect to database due to ", err)
		return
	}
	db.Where(&user).Find(&user)
	if user.ID == 0 {
		//No user
		user.ID = uint(time.Now().Unix())
		db.Create(&user)   //Creating a user
		w.WriteHeader(201) //Created
	} else {
		//User Exists Previously
		w.WriteHeader(400) //Bad Request
	}
}

//login, logins a new user
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vals := r.PostForm
	hpasswd := vals.Get("password")
	username := vals.Get("username")

	user := User{UserName: username, Password: hpasswd}
	db, err := getDataBase()
	if err != nil {
		w.WriteHeader(500)
		log.Print("Couldn't connect to database due to ", err)
		return
	}
	db.Where(&user).Find(&user)
	if user.ID == 0 {
		w.WriteHeader(403) //Forbidden
	} else {

		b, _ := json.Marshal(map[string]interface{}{"id": user.ID})
		v := w.Header()
		v.Set("Content-Type", "application/json")
		w.Write(b)
		w.WriteHeader(200)
	}
}

//changePasswd  changes user password
func changePasswd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vals := r.PostForm
	username := vals.Get("username")
	hpasswd := vals.Get("password")
	newpassword := vals.Get("new_password")
	user := User{UserName: username, Password: hpasswd}
	db, err := getDataBase()
	if err != nil {
		w.WriteHeader(500)
		log.Print("Couldn't connect to database due to ", err)
		return
	}

	db.Where(&user).Find(&user)
	affect := db.Model(&user).Update("password", newpassword).RowsAffected
	if affect == 1 {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403) //Forbidden
	}
}
