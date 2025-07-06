package main

import (
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func users (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users Page"))
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":8086", nil))
}
