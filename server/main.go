package main

import (
	"fmt"
	"log"
	"net/http"
)

type Address struct {
	port string
}

var serverAddress Address = Address{
	port: ":5000",
}

func main() {
	fmt.Println("Server Running On Port 5000")
	http.HandleFunc("/", HandleRequsts)
	log.Fatal(http.ListenAndServe(serverAddress.port, nil))
}
