package main

type User struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

var db []User
