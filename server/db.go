package main

type User struct {
	ID       string `json:"id"`
	NAME     string `json:"name"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}
type Image struct {
	ID       string `json:"id"`
	FILENAME string `json:"filename"`
	USERID   string `json:"userid"`
}

var UserDB []User
var imageDB []Image
