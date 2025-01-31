package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var current_user = make(map[string]string)

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("something went wrong")
		return
	}
	email := r.FormValue("email")

	for _, user := range UserDB {
		if user.EMAIL == email {
			current_user["id"] = string(user.ID)
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
	user := User{NAME: r.FormValue("name"), EMAIL: r.FormValue("email"), PASSWORD: r.FormValue("password"), ID: strconv.Itoa(rand.Int())}
	UserDB = append(UserDB, user)
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
	var userImages []Image
	for _, image := range imageDB {
		if current_user["id"] == image.USERID {
			userImages = append(userImages, image)
		}
	}
	data, err := json.Marshal(userImages)
	if err != nil {
		fmt.Println("failed conversion")
		return
	}
	w.Write(data)
}
func uploadPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../view/upload.html")
}
func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Filename)
	filepath := "./uploads/" + handler.Filename
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	newImage := Image{FILENAME: filepath, ID: strconv.Itoa(rand.Int()), USERID: current_user["id"]}
	imageDB = append(imageDB, newImage)
	http.Redirect(w, r, "/home", http.StatusAccepted)
}
func logout(w http.ResponseWriter, r *http.Request) {
	current_user["id"] = ""
	http.Redirect(w, r, "/login", http.StatusAccepted)
}
func HandleRequsts(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	route := url[len(url)-1]
	if r.Method == "GET" && route == "login" {
		loginPage(w, r)
	} else if r.Method == "GET" && route == "register" {
		registerPage(w, r)
	} else if r.Method == "POST" && route == "login" {
		login(w, r)
	} else if r.Method == "POST" && route == "register" {
		register(w, r)
	}

	if r.Method == "GET" && route == "home" && current_user["id"] != "" {
		homePage(w, r)
	} else if r.Method == "GET" && route == "upload" && current_user["id"] != "" {
		uploadPage(w, r)
	} else if r.Method == "POST" && route == "upload" && current_user["id"] != "" {
		upload(w, r)
	} else {
		http.Redirect(w, r, "/login", http.StatusBadRequest)
	}

}
