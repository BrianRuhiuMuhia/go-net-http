package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	fmt.Println(email)
	for _, user := range db {
		if user.EMAIL == email {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/register", http.StatusSeeOther)

}
func register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("something went wrong with the form")
		w.Write([]byte("something went wrong"))
	}
	if r.FormValue("password") != r.FormValue("confirm-password") {
		w.Write([]byte("passwords must match"))
	}
	user := User{NAME: r.FormValue("name"), EMAIL: r.FormValue("email"), PASSWORD: r.FormValue("password"), ID: rand.Int()}
	db = append(db, user)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func loginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../view/login.html")
}
func registerPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../view/register.html")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(db)
	if err != nil {
		fmt.Println("there was an error")
	}
	w.Write(data)
}
func HandleRequsts(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	route := url[len(url)-1]
	if r.Method == "GET" && route == "login" {
		loginPage(w, r)
	} else if r.Method == "GET" && route == "register" {
		registerPage(w, r)
	} else if r.Method == "GET" && route == "home" {
		homePage(w, r)
	} else if r.Method == "POST" && route == "login" {
		login(w, r)
	} else if r.Method == "POST" && route == "register" {
		register(w, r)
	}

}
